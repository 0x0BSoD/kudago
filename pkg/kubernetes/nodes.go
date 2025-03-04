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

// GetNodesJSON returns the list of nodes in JSON format.
func (c *Client) GetNodesJSON(ctx context.Context) (string, error) {
	nodes, err := c.GetNodes(ctx)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(nodes)
	if err != nil {
		return "", fmt.Errorf("failed to marshal nodes list: %w", err)
	}

	return string(data), nil
}
