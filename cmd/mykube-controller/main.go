package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
		fmt.Println(&kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
			os.Exit(1)
			// panic(err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Cannot make Get call: %v\n", err)
		os.Exit(1)
	}

	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), "default", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Cannot make Get call: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Logging pod name: %v\n", pod.GetName())
	fmt.Printf("Logging ns status: %v\n", ns.Status)

}
