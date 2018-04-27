package main

import (
	"flag"
	"path/filepath"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"fmt"
	"os"
	"k8s.io/api/core/v1"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Namespaces")
	fmt.Println("************************")
	ns, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d NAMESPACES in the cluster.\n", len(ns.Items))

	fmt.Println("Services")
	fmt.Println("*************************")
	svcs, err := clientset.CoreV1().Services("kube-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d Services in the kube-system.\n", len(svcs.Items))


	fmt.Println("Deployments")
	fmt.Println("*************************")
	dps, err := clientset.AppsV1beta1().Deployments("kube-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d Deployments in the kube-system namespace.\n", len(dps.Items))

	fmt.Println("Try create a namespace")
	nsSpec := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name:"officalsite"}}
	_, err = clientset.CoreV1().Namespaces().Create(nsSpec)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Namespaces")
	fmt.Println("************************")
	ns, err = clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d NAMESPACES in the cluster.\n", len(ns.Items))
}

func homeDir() string {
	return os.Getenv("HOME")
}
