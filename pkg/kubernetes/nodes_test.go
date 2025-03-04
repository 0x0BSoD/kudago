package kubernetes

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

// TestGetNodes tests the GetNodes function using a fake Kubernetes client.
func TestGetNodes(t *testing.T) {

	fakeClient := fake.NewSimpleClientset(
		&v1.Node{
			ObjectMeta: metaV1.ObjectMeta{Name: "node1"},
		},
		&v1.Node{
			ObjectMeta: metaV1.ObjectMeta{Name: "node2"},
		})

	client := &Client{
		CoreV1Api: fakeClient.CoreV1(),
	}

	nodes, err := client.GetNodes(context.TODO())
	if err != nil {
		t.Fatalf("GetNodes failed: %v", err)
	}

	if len(nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(nodes))
	}

	if nodes[0].Name != "node1" || nodes[1].Name != "node2" {
		t.Errorf("Unexpected node names: %+v", nodes)
	}
}
