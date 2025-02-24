package kubernetes

import (
	"errors"
	"io/fs"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os"
	"path/filepath"
)

type Client struct {
	ApiClient *kubernetes.Clientset
	Config    *api.Config
}

func New() Client {
	var configPath string

	envVal := os.Getenv("KUBECONFIG")
	if envVal != "" {
		configPath = envVal
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	stdConfigPath := filepath.Join(homedir, ".kube", "config")
	_, err = os.Stat(stdConfigPath)
	if err == nil {
		configPath = stdConfigPath
	}
	if errors.Is(err, fs.ErrNotExist) {
		return Client{}
	}

	// get configs
	apiConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return Client{}
	}

	config, err := clientcmd.LoadFromFile(configPath)
	if err != nil {
		return Client{}
	}

	// create API clinet
	clientSet, err := kubernetes.NewForConfig(apiConfig)
	if err != nil {
		return Client{}
	}

	return Client{
		ApiClient: clientSet,
		Config:    config,
	}
}
