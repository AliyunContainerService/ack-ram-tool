package binding

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"time"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Kind string

var (
	KindRoleBinding        Kind = "RoleBinding"
	KindClusterRoleBinding Kind = "ClusterRoleBinding"
)

type RawBindings struct {
	RoleBindings        []rbacv1.RoleBinding
	ClusterRoleBindings []rbacv1.ClusterRoleBinding
}

var errInvalidName = errors.New("invalid name")
var regexAliUserIdentity = regexp.MustCompile(`^(\d+)(-\d+)?$`)

func (bs *RawBindings) AliUserBindings() RawBindings {
	filtered := RawBindings{}
	for _, b := range bs.RoleBindings {
		if isAliUserRoleBinding(b) {
			filtered.RoleBindings = append(filtered.RoleBindings, b)
		}
	}
	for _, b := range bs.ClusterRoleBindings {
		if isAliUserClusterRoleBinding(b) {
			filtered.ClusterRoleBindings = append(filtered.ClusterRoleBindings, b)
		}
	}
	return filtered
}

type Binding struct {
	Kind        Kind
	Name        string
	Namespace   string
	SubjectName string
	AliUid      int64
}

func (bs *RawBindings) SortByUid() []Binding {
	var bindList []Binding
	for _, b := range bs.RoleBindings {
		for _, sub := range b.Subjects {
			bindList = append(bindList, Binding{
				Kind:        KindRoleBinding,
				Name:        b.Name,
				Namespace:   b.Namespace,
				SubjectName: sub.Name,
				AliUid:      0,
			})
		}
	}
	for _, b := range bs.ClusterRoleBindings {
		for _, sub := range b.Subjects {
			bindList = append(bindList, Binding{
				Kind:        KindClusterRoleBinding,
				Name:        b.Name,
				Namespace:   b.Namespace,
				SubjectName: sub.Name,
				AliUid:      0,
			})
		}
	}
	for i, b := range bindList {
		b := b
		uid, err := getAliUidFromSubjectName(b.SubjectName)
		if err != nil || uid == 0 {
			continue
		}
		b.AliUid = uid
		bindList[i] = b
	}
	sort.Slice(bindList, func(i, j int) bool {
		return bindList[i].AliUid < bindList[j].AliUid
	})
	return bindList
}

func ListBindings(ctx context.Context, kube kubernetes.Interface) (*RawBindings, error) {
	roleBindings, err := listRoleBindings(ctx, kube)
	if err != nil {
		return nil, err
	}
	clusterRoleBindings, err := listClusterRoleBindings(ctx, kube)
	if err != nil {
		return nil, err
	}
	return &RawBindings{
		RoleBindings:        roleBindings.Items,
		ClusterRoleBindings: clusterRoleBindings.Items,
	}, nil
}

func isAliUserRoleBinding(binding rbacv1.RoleBinding) bool {
	for _, sub := range binding.Subjects {
		if isAliUserSubject(sub) {
			return true
		}
	}
	return false
}

func isAliUserClusterRoleBinding(binding rbacv1.ClusterRoleBinding) bool {
	for _, sub := range binding.Subjects {
		if isAliUserSubject(sub) {
			return true
		}
	}
	return false
}

func getAliUidFromSubjectName(name string) (int64, error) {
	matches := regexAliUserIdentity.FindAllStringSubmatch(name, -1)
	if len(matches) < 1 {
		return 0, errInvalidName
	}
	match := matches[0]
	if len(match) < 2 {
		return 0, errInvalidName
	}
	uid, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, err
	}
	return int64(uid), nil
}

func isAliUserSubject(subject rbacv1.Subject) bool {
	if subject.Kind != rbacv1.UserKind {
		return false
	}
	if subject.APIGroup != rbacv1.GroupName {
		return false
	}

	return regexAliUserIdentity.MatchString(subject.Name)
}

func listRoleBindings(ctx context.Context, kube kubernetes.Interface) (*rbacv1.RoleBindingList, error) {
	continueMark := ""
	allList := &rbacv1.RoleBindingList{Items: []rbacv1.RoleBinding{}}
	i := 0
	limit := int64(500)
	maxI := limit * 1000

	for {
		opt := metav1.ListOptions{
			Limit:    limit,
			Continue: continueMark,
			//ResourceVersion: "0",
		}
		ret, err := kube.RbacV1().RoleBindings("").List(ctx, opt)
		if err != nil {
			return nil, err
		}
		allList.Items = append(allList.Items, ret.Items...)
		continueMark = ret.Continue
		if continueMark == "" || i > int(maxI/limit) {
			break
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}

	return allList, nil
}

func listClusterRoleBindings(ctx context.Context, kube kubernetes.Interface) (*rbacv1.ClusterRoleBindingList, error) {
	continueMark := ""
	allList := &rbacv1.ClusterRoleBindingList{Items: []rbacv1.ClusterRoleBinding{}}
	i := 0
	limit := int64(500)
	maxI := limit * 1000

	for {
		opt := metav1.ListOptions{
			Limit:    limit,
			Continue: continueMark,
			//ResourceVersion: "0",
		}
		ret, err := kube.RbacV1().ClusterRoleBindings().List(ctx, opt)
		if err != nil {
			return nil, err
		}
		allList.Items = append(allList.Items, ret.Items...)
		continueMark = ret.Continue
		if continueMark == "" || i > int(maxI/limit) {
			break
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}

	return allList, nil
}

func SaveBindingToFile(ctx context.Context, dir string, b Binding, client kubernetes.Interface) (string, error) {
	var r interface{}
	var err error
	switch b.Kind {
	case KindRoleBinding:
		r, err = getRoleBinding(ctx, b, client)
		if err != nil {
			return "", err
		}
	case KindClusterRoleBinding:
		r, err = getClusterRoleBinding(ctx, b, client)
		if err != nil {
			return "", err
		}
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("mkdir %s: %w", dir, err)
		}
	}
	filePath := path.Join(dir, fmt.Sprintf("%s-%s-%s.json", b.Kind, b.Namespace, b.Name))
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return "", err
	}
	return filePath, nil
}

func RemoveBinding(ctx context.Context, b Binding, client kubernetes.Interface) error {
	switch b.Kind {
	case KindRoleBinding:
		return removeRoleBinding(ctx, b, client)
	case KindClusterRoleBinding:
		return removeClusterRoleBinding(ctx, b, client)
	}
	return nil
}

func getClusterRoleBinding(ctx context.Context, b Binding, client kubernetes.Interface) (*rbacv1.ClusterRoleBinding, error) {
	obj, err := client.RbacV1().ClusterRoleBindings().Get(ctx, b.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	obj.APIVersion = "rbac.authorization.k8s.io/v1"
	obj.Kind = "ClusterRoleBinding"
	return obj, err
}

func getRoleBinding(ctx context.Context, b Binding, client kubernetes.Interface) (*rbacv1.RoleBinding, error) {
	obj, err := client.RbacV1().RoleBindings(b.Namespace).Get(ctx, b.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	obj.APIVersion = "rbac.authorization.k8s.io/v1"
	obj.Kind = "RoleBinding"
	return obj, err
}

func removeClusterRoleBinding(ctx context.Context, b Binding, client kubernetes.Interface) error {
	err := client.RbacV1().ClusterRoleBindings().Delete(ctx, b.Name, metav1.DeleteOptions{})
	return err
}

func removeRoleBinding(ctx context.Context, b Binding, client kubernetes.Interface) error {
	err := client.RbacV1().RoleBindings(b.Namespace).Delete(ctx, b.Name, metav1.DeleteOptions{})
	return err
}

func (b Binding) String() string {
	switch b.Kind {
	case KindRoleBinding:
		return fmt.Sprintf("%s/%s/%s", b.Kind, b.Namespace, b.Name)
	case KindClusterRoleBinding:
		return fmt.Sprintf("%s/-/%s", b.Kind, b.Name)
	}
	return fmt.Sprintf("%#v", b)
}
