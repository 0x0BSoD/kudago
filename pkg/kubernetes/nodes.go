package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNodes retrieves a list of Kubernetes nodes and returns them as a Go slice.
func (c *Client) GetNodes(ctx context.Context) ([]v1.Node, error) {
	if c.CoreV1Api == nil {
		return nil, fmt.Errorf("kubernetes API client is not initialized")
	}

	nodesList, err := c.CoreV1Api.Nodes().List(ctx, metaV1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes: %w", err)
	}

	return nodesList.Items, nil
}

func (c *Client) GetNode(ctx context.Context, name string) (*v1.Node, error) {
	if c.CoreV1Api == nil {
		return nil, fmt.Errorf("kubernetes API client is not initialized")
	}

	node, err := c.CoreV1Api.Nodes().Get(ctx, name, metaV1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes: %w", err)
	}

	return node, nil
}

func (c *Client) GetResourceJSON(ctx context.Context, getResource func() (interface{}, error)) (string, error) {
	resource, err := getResource()
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(resource)
	if err != nil {
		return "", fmt.Errorf("failed to marshal resource: %w", err)
	}

	return string(data), nil
}
