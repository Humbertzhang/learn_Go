package main

import (
	"os"
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	appsv1b1 "k8s.io/api/extensions/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"


	"fmt"
)

func main()  {
	var kubeconfig *string
	if home := myhomeDir(); home != "" {
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

	deploys := clientset.ExtensionsV1beta1().Deployments("fuck")
	//deploymentClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1b1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:"client-go-test",
		},
		Spec:appsv1b1.DeploymentSpec{
			Replicas: Int32Ptr(2),
			/*Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string {
					"app": "demo",
				},
			},*/
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":"demo",
					},
				},
				Spec:apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:"test",
							Image:"nginx",
							Ports:[]apiv1.ContainerPort{
								{
									Name:"http",
									Protocol: apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	result, err := deploys.Create(deployment)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("Create deployment %q. \n", result.GetObjectMeta().GetName())
}

func myhomeDir() string {
	return os.Getenv("HOME")
}

func Int32Ptr(i int32) *int32 {
	return &i
}
