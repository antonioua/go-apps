package main

import (
	"context"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	clientSet := kubernetes.NewForConfigOrDie(config)
}
