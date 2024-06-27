package scanaddon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sort"
	"strings"
)

type ClusterScanner struct {
	openAPIClient openapi.ClientInterface
	kubeClient    kubernetes.Interface

	clusterId string
}

func NewClusterScanner(openAPIClient openapi.ClientInterface,
	kubeClient kubernetes.Interface, clusterId string) *ClusterScanner {
	return &ClusterScanner{
		openAPIClient: openAPIClient,
		kubeClient:    kubeClient,
		clusterId:     clusterId,
	}
}

func (s *ClusterScanner) Scan(ctx context.Context) error {
	needs, err := s.scan(ctx)
	if err != nil {
		return err
	}

	log.Logger.Info("Summary of the Scan Result:")
	if len(needs) == 0 {
		utils.PrintlnToStdErr("● No system addons need to be updated!")
		return nil
	}
	var actions []string
	for name, item := range needs {
		var policyNames []string
		for _, p := range item.PolicyNames() {
			policyNames = append(policyNames, utils.Underline(p))
		}
		msg := fmt.Sprintf("● Addon %s needs to be updated, which depends on %s",
			utils.Underline(name), strings.Join(policyNames, ", "))
		utils.PrintlnToStdErr(msg)
		for _, ac := range item.PolicyActions() {
			if utils.StringSliceInclude(actions, ac) {
				continue
			}
			actions = append(actions, ac)
		}
	}
	sort.Strings(actions)

	actionsJson, _ := json.MarshalIndent(actions, " ", " ")
	utils.PrintlnToStdErr("● These addons need " + utils.Underline("RAM actions") +
		" as follows:")
	fmt.Println(strings.Trim(string(actionsJson), "[]"))

	return nil
}

func (s *ClusterScanner) scan(ctx context.Context) (map[string]NeedUpdateAddon, error) {
	installedAddons, err := s.listAddons(ctx)
	if err != nil {
		return nil, err
	}
	secretList, err := listSecret(ctx, s.kubeClient, "kube-system")
	if err != nil {
		return nil, fmt.Errorf("list secrets from kube-system: %w", err)
	}
	secrets := map[string]corev1.Secret{}
	for _, item := range secretList.Items {
		secrets[item.Name] = item
	}
	workloads, err := s.listWorkloads(ctx, secrets)
	if err != nil {
		return nil, err
	}
	sort.Slice(workloads, func(i, j int) bool {
		return len(workloads[i].MountedNames) > len(workloads[j].MountedNames)
	})

	log.Logger.Debugf("start to check %d workloads", len(workloads))
	needs := map[string]NeedUpdateAddon{}
	for _, wl := range workloads {
		wl := wl
		lgr := log.Logger.Named(fmt.Sprintf("[%s]", wl.String()))
		ctx := log.IntoContext(ctx, lgr)
		addon, roleNames, hardened, err := s.checkWorkload(ctx, wl, secrets, installedAddons)
		if err != nil {
			return nil, err
		}
		addonName := addon.Name
		workloadName := fmt.Sprintf("%s.%s", wl.Name, wl.Namespace)
		if addonName == "" {
			addonName = workloadName
		}
		if hardened {
			lgr.Debugf("%s was hardened", addonName)
		} else {
			lgr.Warnf("%s should be updated, which depends on %v", addonName, roleNames)
			if _, ok := needs[addonName]; !ok {
				if addon.Name == "" {
					addon.Name = addonName
				}
				needs[addonName] = NeedUpdateAddon{
					Addon: addon,
				}
			}
			na := needs[addonName]
			na.Workloads = append(na.Workloads, wl)
			for _, role := range roleNames {
				if !utils.StringSliceInclude(na.RoleNames, role) {
					na.RoleNames = append(na.RoleNames, role)
				}
			}
			needs[addonName] = na
		}
	}

	return needs, nil
}

func (s *ClusterScanner) checkWorkload(ctx context.Context, wl Workload,
	secrets map[string]corev1.Secret, installedAddons map[string]InstalledAddon) (
	addon Addon, roleNames []string, hardened bool, err error) {
	logger := log.FromContext(ctx)
	logger.Debugf("mounted: %v, Hardened: %v, Images: %v",
		wl.MountedNames, wl.Hardened, wl.ImageNames)

	addon, ok := getAddonByWorkload(wl, installedAddons)
	if ok {
		logger.Debugf("match addon %s", addon.Name)
		if len(wl.MountedNames) == 0 {
			logger.Debugf("no mount names with addon %s", addon.Name)
		}
	} else {
		logger.Debugf("no addon match,"+
			"mounted: %v, %v, Images: %v",
			wl.MountedNames, wl.Hardened, wl.ImageNames)
	}

	if err := s.prepareWorkload(ctx, addon, &wl, secrets, installedAddons); err != nil {
		return addon, nil, false, err
	}
	if len(wl.MountedNames) > 0 && len(wl.MountedNames) == len(wl.Hardened) {
		return addon, nil, true, nil
	}

	var mountedRoles []string
	prefix := "addon."
	suffix := ".token"
	for k, v := range roleToName {
		for _, name := range wl.MountedNames {
			if utils.StringSliceInclude(wl.Hardened, name) {
				continue
			}
			if prefix+v+suffix == name {
				mountedRoles = append(mountedRoles, k)
			}
		}
	}

	if len(mountedRoles) == 0 {
		if len(addon.DefaultRoleNames) > 0 {
			mountedRoles = append(mountedRoles, addon.DefaultRoleNames...)
		} else {
			logger.Debugf("not match addon and no mount, skip it")
			return addon, nil, true, nil
		}
	}
	return addon, mountedRoles, false, nil
}

func (s *ClusterScanner) prepareWorkload(ctx context.Context,
	addon Addon, wl *Workload, secrets map[string]corev1.Secret,
	installedAddons map[string]InstalledAddon) error {
	if err := s.prepareTerway(ctx, addon, wl); err != nil {
		return err
	}
	if err := s.prepareMseController(ctx, addon, wl, secrets, installedAddons); err != nil {
		return err
	}
	if err := s.prepareArmsPrometheus(ctx, addon, wl, secrets, installedAddons); err != nil {
		return err
	}
	if err := s.prepareArmsCmonitor(ctx, addon, wl, secrets, installedAddons); err != nil {
		return err
	}
	if err := s.prepareOnepilot(ctx, addon, wl, secrets, installedAddons); err != nil {
		return err
	}
	if err := s.prepareKubeAi(ctx, addon, wl, secrets, installedAddons); err != nil {
		return err
	}

	return nil
}

func (s *ClusterScanner) prepareMseController(ctx context.Context,
	addon Addon, wl *Workload, secrets map[string]corev1.Secret,
	installedAddons map[string]InstalledAddon) error {
	logger := log.FromContext(ctx)
	switch addon.Name {
	case "mse-ingress-controller":
		break
	default:
		return nil
	}
	if len(wl.MountedNames) > 0 {
		return nil
	}
	keyword := "aliyuncsmanagedmserole"
	prefix := "addon."
	suffix := ".token"
	wl.MountedNames = []string{prefix + keyword + suffix}

	role, err := s.getRole(ctx, "kube-system", "ack-mse-ingress-controller-addon")
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Debugf("get role failed: %s", err)
			return nil
		}
		return err
	}

	if !checkRoleResourcesIncludeKeyword(role, keyword) {
		logger.Debugf("rules of role %s not include %s", role.Name, keyword)
		return nil
	}

	wl.Hardened = analyzeMount(ctx, secrets, wl.MountedNames, wl.CreateTime, wl.ReadyTime)

	return nil
}

func (s *ClusterScanner) prepareArmsPrometheus(ctx context.Context,
	addon Addon, wl *Workload, secrets map[string]corev1.Secret,
	installedAddons map[string]InstalledAddon) error {
	logger := log.FromContext(ctx)
	switch addon.Name {
	case "arms-prometheus":
		break
	default:
		return nil
	}
	if len(wl.MountedNames) > 0 {
		return nil
	}

	if wl.Name != "o11y-addon-controller" {
		return nil
	}

	keyword := "arms"
	prefix := "addon."
	suffix := ".token"
	wl.MountedNames = []string{prefix + keyword + suffix}

	role, err := s.getClusterRole(ctx, "o11y:addon-controller:role")
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Debugf("get role failed: %s", err)
			return nil
		}
		return err
	}

	if !checkClusterRoleResourcesIncludeKeyword(role, keyword) {
		logger.Debugf("rules of clusterrole %s not include %s", role.Name, keyword)
		return nil
	}

	wl.Hardened = analyzeMount(ctx, secrets, wl.MountedNames, wl.CreateTime, wl.ReadyTime)

	return nil
}

func (s *ClusterScanner) prepareArmsCmonitor(ctx context.Context,
	addon Addon, wl *Workload, secrets map[string]corev1.Secret,
	installedAddons map[string]InstalledAddon) error {
	logger := log.FromContext(ctx)
	switch addon.Name {
	case "arms-cmonitor":
		break
	default:
		return nil
	}
	if len(wl.MountedNames) > 0 {
		return nil
	}

	prefix := "addon."
	suffix := ".token"
	allExist := true
	for _, item := range []string{"arms,ot-collector-cluster-role", "aliyuncsmanagedmserole,ot-collector-cluster-role"} {
		parts := strings.Split(item, ",")
		keyword := parts[0]
		roleName := parts[1]
		wl.MountedNames = append(wl.MountedNames, prefix+keyword+suffix)

		role, err := s.getClusterRole(ctx, roleName)
		if err != nil {
			if errors.IsNotFound(err) {
				logger.Debugf("get role failed: %s", err)
				allExist = false
				continue
			}
			return err
		}
		if !checkClusterRoleResourcesIncludeKeyword(role, keyword) {
			logger.Debugf("rules of clusterrole %s not include %s", role.Name, keyword)
			return nil
		}
		if _, ok := secrets[prefix+keyword+suffix]; !ok {
			allExist = false
		}
	}

	wl.Hardened = analyzeMount(ctx, secrets, wl.MountedNames, wl.CreateTime, wl.ReadyTime)
	if allExist && utils.StringSliceInclude(wl.Hardened, prefix+"arms"+suffix) {
		wl.Hardened = wl.MountedNames
	}
	return nil
}

func (s *ClusterScanner) prepareOnepilot(ctx context.Context,
	addon Addon, wl *Workload, secrets map[string]corev1.Secret,
	installedAddons map[string]InstalledAddon) error {
	logger := log.FromContext(ctx)
	switch addon.Name {
	case "ack-onepilot":
		break
	default:
		return nil
	}
	if len(wl.MountedNames) > 0 {
		return nil
	}

	prefix := "addon."
	suffix := ".token"
	allExist := true
	for _, item := range []string{"arms,ack-onepilot-ack-onepilot-role", "aliyuncsmanagedmserole,ack-onepilot-ack-onepilot-role"} {
		parts := strings.Split(item, ",")
		keyword := parts[0]
		roleName := parts[1]
		wl.MountedNames = append(wl.MountedNames, prefix+keyword+suffix)

		role, err := s.getRole(ctx, "kube-system", roleName)
		if err != nil {
			if errors.IsNotFound(err) {
				logger.Debugf("get role failed: %s", err)
				allExist = false
				continue
			}
			return err
		}
		if !checkRoleResourcesIncludeKeyword(role, keyword) {
			logger.Debugf("rules of clusterrole %s not include %s", role.Name, keyword)
			return nil
		}
		if _, ok := secrets[prefix+keyword+suffix]; !ok {
			allExist = false
		}
	}

	wl.Hardened = analyzeMount(ctx, secrets, wl.MountedNames, wl.CreateTime, wl.ReadyTime)
	if allExist && utils.StringSliceInclude(wl.Hardened, prefix+"arms"+suffix) {
		wl.Hardened = wl.MountedNames
	}
	return nil
}

func (s *ClusterScanner) prepareKubeAi(ctx context.Context,
	addon Addon, wl *Workload, secrets map[string]corev1.Secret,
	installedAddons map[string]InstalledAddon) error {
	switch addon.Name {
	case "kube-ai":
		break
	default:
		return nil
	}
	if len(wl.MountedNames) > 0 {
		return nil
	}

	prefix := "addon."
	suffix := ".token"

	wl.MountedNames = append(wl.MountedNames, prefix+"kubeai"+suffix)

	return nil
}

func checkRoleResourcesIncludeKeyword(role *rbacv1.Role, keyword string) bool {
	for _, rule := range role.Rules {
		if utils.StringSliceInclude(rule.Resources, "secrets") {
			for _, name := range rule.ResourceNames {
				if strings.Contains(name, "."+keyword+".") {
					return true
				}
			}
		}
	}
	return false
}

func checkClusterRoleResourcesIncludeKeyword(role *rbacv1.ClusterRole, keyword string) bool {
	for _, rule := range role.Rules {
		if utils.StringSliceInclude(rule.Resources, "secrets") {
			for _, name := range rule.ResourceNames {
				if strings.Contains(name, "."+keyword+".") {
					return true
				}
			}
		}
	}
	return false
}

func (s *ClusterScanner) getRole(ctx context.Context, namespace, name string) (*rbacv1.Role, error) {
	role, err := s.kubeClient.RbacV1().Roles(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get role %s from %s: %w", name, namespace, err)
	}
	return role, nil
}

func (s *ClusterScanner) getClusterRole(ctx context.Context, name string) (*rbacv1.ClusterRole, error) {
	role, err := s.kubeClient.RbacV1().ClusterRoles().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get clusterrole %s: %w", name, err)
	}
	return role, nil
}

func (s *ClusterScanner) prepareTerway(ctx context.Context, addon Addon, wl *Workload) error {
	switch addon.Name {
	case "terway", "terway-eni", "terway-eniip", "terway-controlplane":
		break
	default:
		return nil
	}
	if len(wl.Hardened) == 0 {
		return nil
	}

	cmp, err := s.kubeClient.CoreV1().ConfigMaps("kube-system").Get(ctx,
		"eni-config", metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("get configmap eni-config from kube-system: %w", err)
	}
	if strings.Contains(cmp.Data["eni_conf"], "credential_path") &&
		strings.Contains(cmp.Data["eni_conf"], "/var/addon/token-config") {
		return nil
	}

	log.Logger.Debugf("eni-config not include credential_path, mark %s as need update",
		wl.Name)
	wl.Hardened = []string{}

	return nil
}

func (s *ClusterScanner) listAddons(ctx context.Context) (map[string]InstalledAddon, error) {
	defer newSpinner()()

	addons, err := s.openAPIClient.ListAddons(ctx, s.clusterId)
	if err != nil {
		return nil, err
	}

	installedAddons := map[string]InstalledAddon{}
	for _, addon := range addons {
		addon := addon
		if !addon.Installed() {
			continue
		}
		installedAddons[addon.Name] = InstalledAddon{
			Name:           addon.Name,
			CurrentVersion: addon.Version,
		}
	}

	return installedAddons, nil
}

func (s *ClusterScanner) listWorkloads(ctx context.Context,
	secrets map[string]corev1.Secret) ([]Workload, error) {
	defer newSpinner()()

	var workloads []Workload
	for _, ns := range scanNamespaces {
		log.Logger.Debugf("start to list workloads in namespace %s", ns)
		ws, err := listWorkload(ctx, s.kubeClient, ns, secrets)
		if err != nil {
			return nil, fmt.Errorf("scan workload from %s: %w", ns, err)
		}
		workloads = append(workloads, ws...)
	}

	return workloads, nil
}
