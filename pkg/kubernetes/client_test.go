package kubernetes

import (
	"os"
	"path/filepath"
	"testing"
)

// createTempKubeConfig creates a temporary kubeconfig file and returns its path.
func createTempKubeConfig(t *testing.T, content string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "kubeconfig-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp kubeconfig: %v", err)
	}
	_, err = tmpFile.Write([]byte(content))
	if err != nil {
		t.Fatalf("failed to write to temp kubeconfig: %v", err)
	}
	tmpFile.Close()
	return tmpFile.Name()
}

func TestNewClient(t *testing.T) {
	validKubeConfig := `
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://mock-server
  name: mock-cluster
contexts:
- context:
    cluster: mock-cluster
    user: mock-user
  name: mock-context
current-context: mock-context
users:
- name: mock-user
  user:
    token: mock-token
`

	tests := []struct {
		name            string
		kubeConfigEnv   string
		setupHomeConfig bool
		expectError     bool
	}{
		{
			name:          "Valid KUBECONFIG env var",
			kubeConfigEnv: createTempKubeConfig(t, validKubeConfig),
			expectError:   false,
		},
		{
			name:            "Valid default ~/.kube/config",
			kubeConfigEnv:   "",
			setupHomeConfig: true,
			expectError:     false,
		},
		{
			name:          "Missing kubeconfig",
			kubeConfigEnv: "",
			expectError:   true,
		},
	}

	// Store original env vars
	originalKubeConfig := os.Getenv("KUBECONFIG")
	homeDir, _ := os.UserHomeDir()
	backupConfigPath := filepath.Join(homeDir, ".kube", "config")

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set KUBECONFIG env variable if needed
			if tc.kubeConfigEnv != "" {
				os.Setenv("KUBECONFIG", tc.kubeConfigEnv)
				defer func() {
					err := os.Unsetenv("KUBECONFIG")
					if err != nil {
						t.Fatalf("Failed to unset KUBECONFIG: %v", err)
					}
					os.Setenv("KUBECONFIG", originalKubeConfig)
				}() // Cleanup after test
			} else {
				os.Unsetenv("KUBECONFIG")
			}

			// Mock ~/.kube/config if required
			if tc.setupHomeConfig {
				_ = os.MkdirAll(filepath.Dir(backupConfigPath), 0755)
				_ = os.WriteFile(backupConfigPath, []byte(validKubeConfig), 0644)
				defer func(name string) {
					err := os.Remove(name)
					if err != nil {
						t.Errorf("Failed to remove backup config file: %v", err)
					}
				}(backupConfigPath)
			}

			client, err := New()

			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectError, err)
			}
			if err == nil && client.ApiClient == nil {
				t.Errorf("Expected ApiClient to be initialized, but got nil")
			}
		})
	}
}
