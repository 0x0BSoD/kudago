package kubernetes

import (
	"fmt"
)

func (k *Client) GetContexts() {
	for _, v := range k.Config.Contexts {
		fmt.Printf("Cluster %v\n", v.Cluster)
		fmt.Printf("Namespace %v\n", v.Namespace)
		fmt.Printf("LocationOfOrigin %v\n", v.LocationOfOrigin)
		fmt.Printf("AuthInfo %v\n", v.AuthInfo)
		fmt.Printf("Extensions %v\n\n", v.Extensions)
	}
}
