package main

import (
	"context"
	"fmt"
	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	"github.com/projectcalico/api/pkg/apis/projectcalico/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	//var kubeConfig *string
	//kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err.Error())
	}

	cs, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	_ = v3.AddToScheme(scheme.Scheme)

	schemeGroupVersion := schema.GroupVersion{
		Group:   "crd.projectcalico.org",
		Version: "v1",
	}

	scheme.Scheme.AddKnownTypes(schemeGroupVersion,
		&v3.GlobalNetworkPolicy{},
		&v3.GlobalNetworkPolicyList{},
	)

	metav1.AddToGroupVersion(scheme.Scheme, schemeGroupVersion)

	list, err := cs.ProjectcalicoV3().GlobalNetworkPolicies().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, nwp := range list.Items {
		fmt.Printf("%#v\n", nwp)
	}
}
