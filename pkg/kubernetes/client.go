package kubernetes

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	CoreV1Api corev1.CoreV1Interface
	AppsV1Api appsv1.AppsV1Interface
	Config    *api.Config
}

// getKubeConfigPath determines the best kubeconfig path to use.
func getKubeConfigPath() (string, error) {
	if envVal := os.Getenv("KUBECONFIG"); envVal != "" {
		return envVal, nil
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	stdConfigPath := filepath.Join(homedir, ".kube", "config")
	if _, err := os.Stat(stdConfigPath); err == nil {
		return stdConfigPath, nil
	} else if errors.Is(err, fs.ErrNotExist) {
		return "", fmt.Errorf("kubeconfig file not found at %s", stdConfigPath)
	} else {
		return "", fmt.Errorf("error checking kubeconfig file: %w", err)
	}
}

// New initializes a new Kubernetes client.
func New() (*Client, error) {
	configPath, err := getKubeConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to determine kubeconfig path: %w", err)
	}

	apiConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build API config: %w", err)
	}

	config, err := clientcmd.LoadFromFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig file: %w", err)
	}

	clientSet, err := kubernetes.NewForConfig(apiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes clientset: %w", err)
	}

	coreClient := clientSet.CoreV1()
	appsClient := clientSet.AppsV1()

	return &Client{
		AppsV1Api: appsClient,
		CoreV1Api: coreClient,
		Config:    config,
	}, nil
}
