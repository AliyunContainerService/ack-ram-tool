package scanuserpermissions

import (
	"context"
	"errors"
	"regexp"
	"sort"
	"strconv"
	"time"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type BindingKind string

var (
	BindingKindRoleBinding        BindingKind = "RoleBinding"
	BindingKindClusterRoleBinding BindingKind = "ClusterRoleBinding"
)

type RawBindings struct {
	RoleBindings        []rbacv1.RoleBinding
	ClusterRoleBindings []rbacv1.ClusterRoleBinding
}

var errInvalidName = errors.New("invalid name")

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
	Kind        BindingKind
	Name        string
	SubjectName string
	AliUid      int
}

func (bs *RawBindings) SortByUid() []Binding {
	var bindList []Binding
	for _, b := range bs.RoleBindings {
		for _, sub := range b.Subjects {
			bindList = append(bindList, Binding{
				Kind:        BindingKindRoleBinding,
				Name:        b.Name,
				SubjectName: sub.Name,
				AliUid:      0,
			})
		}
	}
	for _, b := range bs.ClusterRoleBindings {
		for _, sub := range b.Subjects {
			bindList = append(bindList, Binding{
				Kind:        BindingKindClusterRoleBinding,
				Name:        b.Name,
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

func listBindings(ctx context.Context, kube kubernetes.Interface) (*RawBindings, error) {
	rolebindings, err := listRoleBindings(ctx, kube)
	if err != nil {
		return nil, err
	}
	clusterroleBindings, err := listClusterRoleBindings(ctx, kube)
	if err != nil {
		return nil, err
	}
	return &RawBindings{
		RoleBindings:        rolebindings.Items,
		ClusterRoleBindings: clusterroleBindings.Items,
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

func getAliUidFromSubjectName(name string) (int, error) {
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
	return uid, nil
}

var regexAliUserIdentity = regexp.MustCompile(`^(\d+)(-\d+)?$`)

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
