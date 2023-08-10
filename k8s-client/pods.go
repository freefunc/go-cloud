package k8s_client

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GetPodCommand struct {
	BaseCommand
	PodName string
}

func (c *GetPodCommand) Execute(cli *KubeClient) (any, error) {
	pod, err := cli.clientset.CoreV1().Pods(c.Namespace).Get(context.TODO(), c.PodName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

type UpdatePodCommand struct {
	BaseCommand
	pod *corev1.Pod
}

func (c *UpdatePodCommand) Execute(cli *KubeClient) (any, error) {
	pod, err := cli.clientset.CoreV1().Pods(c.Namespace).Update(context.TODO(), c.pod, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}
