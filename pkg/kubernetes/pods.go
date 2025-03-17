package kubernetes

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetAllPods(ctx context.Context) ([]v1.Pod, error) {
	if c.CoreV1Api == nil {
		return nil, fmt.Errorf("kubernetes API client is not initialized")
	}

	pods, err := c.CoreV1Api.Pods("").List(ctx, metaV1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pods from namespace: %w", err)
	}

	return pods.Items, nil
}

func (c *Client) GetPod(ctx context.Context, ns string, name string) (*v1.Pod, error) {
	if c.CoreV1Api == nil {
		return nil, fmt.Errorf("kubernetes API client is not initialized")
	}

	pod, err := c.CoreV1Api.Pods(ns).Get(ctx, name, metaV1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pod from namespace %s: %w", ns, err)
	}

	return pod, nil
}
