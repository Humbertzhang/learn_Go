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
	"k8s.io/apimachinery/pkg/util/intstr"
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

	deploys := clientset.ExtensionsV1beta1().Deployments("officalsite")
	//deploymentClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1b1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:"muxistudio-offical-site",
		},
		Spec:appsv1b1.DeploymentSpec{
			Replicas: Int32Ptr(1),
			/*Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string {
					"app": "demo",
				},
			},*/
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:"muxistudio",
					Namespace:"officalsite",
					Labels: map[string]string{
						"run":"muxistudio",
					},
				},
				Spec:apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:"muxistudio",
							Image:"registry.cn-shenzhen.aliyuncs.com/muxi_site/muxi-studio:3.0-beta2",
							Ports:[]apiv1.ContainerPort{
								{
									Name:"http",
									Protocol: apiv1.ProtocolTCP,
									ContainerPort: 3000,

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
	fmt.Printf("Create deployment %q. \n", result.GetObjectMeta().GetName())

	svcs := clientset.CoreV1().Services("officalsite")
	svc := &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:"Service",
			APIVersion:"v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:"muxistudio",
			Namespace: "officalsite",
			Labels: map[string] string{
				"run":"muxistudio",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Protocol:apiv1.ProtocolTCP,
					Port: 3000,
					TargetPort: intstr.IntOrString{IntVal:3000, StrVal:"3000"},
				},
			},
			Selector: map[string]string{
				"run":"muxistudio",
			},
		},
	}
	fmt.Println("Creating Service....")
	resultsvc, err := svcs.Create(svc)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Creating svc %q. \n", resultsvc.GetObjectMeta().GetName())

}

func myhomeDir() string {
	return os.Getenv("HOME")
}

func Int32Ptr(i int32) *int32 {
	return &i
}

