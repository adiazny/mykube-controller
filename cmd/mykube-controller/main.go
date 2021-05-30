package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		kubeconfig := filepath.Join("~", ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		fmt.Println(kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
			os.Exit(1)
			// panic(err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "example", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Cannot make Get call: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Logging pod name: %v", pod.GetName())

}
