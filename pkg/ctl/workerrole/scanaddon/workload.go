package scanaddon

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"strings"
	"time"
)

func listWorkload(ctx context.Context, client kubernetes.Interface,
	namespace string, secrets map[string]corev1.Secret) ([]Workload, error) {
	logger := log.FromContext(ctx)
	var workloads []Workload

	deploymentList, err := listDeployment(ctx, client, namespace)
	if err != nil {
		logger.Errorf("list deployment failed: %s", err)
		return nil, err
	}
	daemonsetList, err := listDaemonSet(ctx, client, namespace)
	if err != nil {
		logger.Errorf("list deamonset failed: %s", err)
		return nil, err
	}

	for _, deployment := range deploymentList.Items {
		lgr := logger.Named(fmt.Sprintf("[%s.%s]", deployment.Name, deployment.Namespace))
		ctx := log.IntoContext(ctx, lgr)
		w, err := parseDeployment(ctx, client, deployment, secrets)
		if err != nil {
			return nil, err
		}
		workloads = append(workloads, *w)
	}

	for _, daemonSet := range daemonsetList.Items {
		lgr := logger.Named(fmt.Sprintf("[%s.%s]", daemonSet.Name, daemonSet.Namespace))
		ctx := log.IntoContext(ctx, lgr)
		w, err := parseDaemonSet(ctx, client, daemonSet, secrets)
		if err != nil {
			return nil, err
		}
		workloads = append(workloads, *w)
	}

	return workloads, nil
}

func parseDaemonSet(ctx context.Context, client kubernetes.Interface,
	daemonSet appsv1.DaemonSet, secrets map[string]corev1.Secret) (*Workload, error) {
	logger := log.FromContext(ctx)
	w := &Workload{
		Type:         WorkloadTypeDaemonSet,
		Namespace:    daemonSet.Namespace,
		Name:         daemonSet.Name,
		Images:       []string{},
		MountedNames: []string{},
		Hardened:     []string{},
	}

	logger.Debug("start check daemonSet")
	for _, cs := range daemonSet.Spec.Template.Spec.Containers {
		parts := strings.Split(cs.Image, "/")
		image := parts[len(parts)-1]
		if !utils.StringSliceInclude(w.Images, image) {
			w.Images = append(w.Images, image)
			w.ImageNames = append(w.ImageNames, strings.Split(image, ":")[0])
			logger.Debugf("new image: %s", image)
		}
	}
	for _, vs := range daemonSet.Spec.Template.Spec.Volumes {
		if vs.Secret == nil {
			continue
		}
		if strings.HasPrefix(vs.Secret.SecretName, "addon.") &&
			strings.HasSuffix(vs.Secret.SecretName, ".token") {
			w.MountedNames = append(w.MountedNames, vs.Secret.SecretName)
			logger.Debugf("got %s", vs.Secret.SecretName)
		}
	}

	readyTime, err := getDaemonSetReadyTime(ctx, client, daemonSet)
	if err != nil {
		return nil, err
	}

	w.CreateTime = daemonSet.CreationTimestamp
	w.ReadyTime = readyTime
	w.Hardened = analyzeMount(ctx, secrets, w.MountedNames,
		daemonSet.CreationTimestamp, readyTime)

	return w, nil
}

func getDaemonSetReadyTime(ctx context.Context, client kubernetes.Interface,
	daemonSet appsv1.DaemonSet) (metav1.Time, error) {
	logger := log.FromContext(ctx)
	ls := labels.Set(daemonSet.Spec.Template.Labels)
	podList, err := listPods(ctx, client, daemonSet.Namespace, ls.String())
	if err != nil {
		logger.Errorf("list pods with labels (%s) failed: %s", ls, err)
		return metav1.Time{}, err
	}
	if len(podList.Items) == 0 {
		logger.Debugf("not found pods with labels: %s", ls)
		return metav1.NewTime(time.Now()), nil
	}

	oldReadyTime := metav1.NewTime(time.Now())
	for _, pod := range podList.Items {
		if pod.Status.Phase != "Running" {
			continue
		}
		for _, cs := range pod.Status.Conditions {
			if cs.Type == "Ready" && cs.Status == "True" {
				rt := cs.LastTransitionTime
				if rt.Before(&oldReadyTime) {
					oldReadyTime = rt
					logger.Debugf("update oldReadyTime to %s", rt.Format(time.RFC3339))
				}
			}
		}
	}

	return oldReadyTime, nil
}

func parseDeployment(ctx context.Context, client kubernetes.Interface,
	deploy appsv1.Deployment, secrets map[string]corev1.Secret) (*Workload, error) {
	logger := log.FromContext(ctx)
	w := &Workload{
		Type:         WorkloadTypeDeployment,
		Namespace:    deploy.Namespace,
		Name:         deploy.Name,
		Images:       []string{},
		MountedNames: []string{},
		Hardened:     []string{},
	}

	logger.Debug("start check deployment")
	for _, cs := range deploy.Spec.Template.Spec.Containers {
		parts := strings.Split(cs.Image, "/")
		image := parts[len(parts)-1]
		if !utils.StringSliceInclude(w.Images, image) {
			w.Images = append(w.Images, image)
			w.ImageNames = append(w.ImageNames, strings.Split(image, ":")[0])
			logger.Debugf("new image: %s", image)
		}
	}
	for _, vs := range deploy.Spec.Template.Spec.Volumes {
		if vs.Secret == nil {
			continue
		}
		if strings.HasPrefix(vs.Secret.SecretName, "addon.") &&
			strings.HasSuffix(vs.Secret.SecretName, ".token") {
			w.MountedNames = append(w.MountedNames, vs.Secret.SecretName)
			logger.Debugf("got %s", vs.Secret.SecretName)
		}
	}

	readyTime, err := getDeploymentReadyTime(ctx, client, deploy)
	if err != nil {
		return nil, err
	}

	w.CreateTime = deploy.CreationTimestamp
	w.ReadyTime = readyTime
	w.Hardened = analyzeMount(ctx, secrets, w.MountedNames,
		deploy.CreationTimestamp, readyTime)

	return w, nil
}

func getDeploymentReadyTime(ctx context.Context, client kubernetes.Interface,
	deploy appsv1.Deployment) (metav1.Time, error) {
	logger := log.FromContext(ctx)
	ls := labels.Set(deploy.Spec.Template.Labels)
	podList, err := listPods(ctx, client, deploy.Namespace, ls.String())
	if err != nil {
		logger.Errorf("list pods with labels (%s) failed: %s", ls, err)
		return metav1.Time{}, err
	}
	if len(podList.Items) == 0 {
		logger.Debugf("not found pods with labels: %s", ls)
		return metav1.NewTime(time.Now()), nil
	}

	oldReadyTime := metav1.NewTime(time.Now())
	for _, pod := range podList.Items {
		if pod.Status.Phase != "Running" {
			continue
		}
		for _, cs := range pod.Status.Conditions {
			if cs.Type == "Ready" && cs.Status == "True" {
				rt := cs.LastTransitionTime
				if rt.Before(&oldReadyTime) {
					oldReadyTime = rt
					logger.Debugf("update oldReadyTime to %s", rt.Format(time.RFC3339))
				}
			}
		}
	}

	return oldReadyTime, nil
}

func analyzeMount(ctx context.Context, secrets map[string]corev1.Secret,
	mountNames []string, createTime, readyTime metav1.Time) []string {
	logger := log.FromContext(ctx)
	var hardened []string

	if readyTime.IsZero() {
		logger.Debug("workload is not ready")
		return hardened
	}

	for _, mn := range mountNames {
		secret, ok := secrets[mn]
		if !ok {
			logger.Debugf("secret not found: %s", mn)
			continue
		}
		if readyTime.Sub(secret.CreationTimestamp.Time) > 0 {
			hardened = append(hardened, mn)
			logger.Debugf("%s, ready time (%s) > secret time (%s)",
				mn, readyTime.Format(time.RFC3339), secret.CreationTimestamp.Format(time.RFC3339))
		} else {
			logger.Debugf("%s, ready time (%s) < secret time (%s)",
				mn, readyTime.Format(time.RFC3339), secret.CreationTimestamp.Format(time.RFC3339))
		}
	}

	return hardened
}

var roleToName = map[string]string{
	"AliyunCSManagedCmsRole":             "cms",
	"AliyunCSManagedLogRole":             "log",
	"AliyunCSManagedNetworkRole":         "network",
	"AliyunCSManagedCsiRole":             "csi",
	"AliyunCSManagedAcrRole":             "aliyuncsmanagedacrrole",
	"AliyunCSManagedCostRole":            "aliyuncsmanagedcostrole",
	"AliyunCSManagedMseRole":             "aliyuncsmanagedmserole",
	"AliyunCSManagedArmsRole":            "arms",
	"AliyunCSManagedAutoScalerRole":      "aliyuncsmanagedautoscalerrole",
	"AliyunCSManagedBackupRestoreRole":   "aliyuncsmanagedbackuprestorerole",
	"AliyunCSManagedWebhookInjectorRole": "aliyuncsmanagedwebhookinjectorrole",
}
