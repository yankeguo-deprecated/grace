package gracek8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// DefaultClient create kubernetes client automatically,
// first try in-cluster client, then from default kubeconfig locations
func DefaultClient() (client *kubernetes.Clientset, err error) {
	// try in-cluster client
	if client, err = InClusterClient(); err == nil {
		return
	}

	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		nil,
	)

	var restCfg *rest.Config
	if restCfg, err = cfg.ClientConfig(); err != nil {
		return
	}

	return kubernetes.NewForConfig(restCfg)
}
