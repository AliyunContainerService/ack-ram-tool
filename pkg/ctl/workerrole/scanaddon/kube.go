package scanaddon

import (
	"context"
	"strings"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

func getKubeClient(ctx context.Context, openAPIClient openapi.ClientInterface,
	clusterId string) (kubernetes.Interface, error) {
	kubeconfig, err := openAPIClient.GetUserKubeConfig(ctx, clusterId,
		opts.privateIpAddress, opts.temporaryDuration)
	if err != nil {
		return nil, err
	}

	client, err := ctlcommon.NewKubeClient(kubeconfig.RawData)
	return client, err
}

func listPods(ctx context.Context, client kubernetes.Interface, namespace string, labels string) (*corev1.PodList, error) {
	continueMark := ""
	allList := &corev1.PodList{Items: []corev1.Pod{}}
	limit := int64(200)

	for {
		opt := metav1.ListOptions{
			Limit:           limit,
			Continue:        continueMark,
			LabelSelector:   labels,
			ResourceVersion: "0",
		}
		ret, err := client.CoreV1().Pods(namespace).List(ctx, opt)
		if err != nil {
			return nil, err
		}
		allList.Items = append(allList.Items, ret.Items...)
		continueMark = ret.Continue
		if continueMark == "" {
			break
		}
		time.Sleep(time.Millisecond * 5)
	}

	return allList, nil
}

func listSecret(ctx context.Context, client kubernetes.Interface, namespace string) (*corev1.SecretList, error) {
	continueMark := ""
	allList := &corev1.SecretList{Items: []corev1.Secret{}}
	limit := int64(200)

	for {
		opt := metav1.ListOptions{
			Limit:           limit,
			Continue:        continueMark,
			ResourceVersion: "0",
		}
		ret, err := client.CoreV1().Secrets(namespace).List(ctx, opt)
		if err != nil {
			return nil, err
		}
		for _, item := range ret.Items {
			if strings.HasPrefix(item.Name, "addon.") &&
				strings.HasSuffix(item.Name, ".token") {
				allList.Items = append(allList.Items, item)
			}
		}
		continueMark = ret.Continue
		if continueMark == "" {
			break
		}
		time.Sleep(time.Millisecond * 5)
	}

	return allList, nil
}

func listDeployment(ctx context.Context, client kubernetes.Interface, namespace string) (*appsv1.DeploymentList, error) {
	continueMark := ""
	allList := &appsv1.DeploymentList{Items: []appsv1.Deployment{}}
	limit := int64(200)

	for {
		opt := metav1.ListOptions{
			Limit:           limit,
			Continue:        continueMark,
			ResourceVersion: "0",
		}
		ret, err := client.AppsV1().Deployments(namespace).List(ctx, opt)
		if err != nil {
			return nil, err
		}
		allList.Items = append(allList.Items, ret.Items...)
		continueMark = ret.Continue
		if continueMark == "" {
			break
		}
		time.Sleep(time.Millisecond * 5)
	}

	return allList, nil
}

func listDaemonSet(ctx context.Context, client kubernetes.Interface, namespace string) (*appsv1.DaemonSetList, error) {
	continueMark := ""
	allList := &appsv1.DaemonSetList{Items: []appsv1.DaemonSet{}}
	limit := int64(200)

	for {
		opt := metav1.ListOptions{
			Limit:           limit,
			Continue:        continueMark,
			ResourceVersion: "0",
		}
		ret, err := client.AppsV1().DaemonSets(namespace).List(ctx, opt)
		if err != nil {
			return nil, err
		}
		allList.Items = append(allList.Items, ret.Items...)
		continueMark = ret.Continue
		if continueMark == "" {
			break
		}
		time.Sleep(time.Millisecond * 5)
	}

	return allList, nil
}
