package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("------------------- Pods -------------------")
		fmt.Println()
		fmt.Println("******************* List All ************************")
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		fmt.Println()
		for _, pod := range pods.Items {
			fmt.Printf("| %6s | %6s |\n", pod.GetName(), pod.GetNamespace())
		}
		fmt.Println()

		fmt.Println("******************* Get By Name *********************")
		podName := "etcd-minikube"
		namespace := "kube-system"
		pod1, _ := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
		fmt.Printf("Found pod %s in namespace %s created at %v\n", podName, namespace, pod1.CreationTimestamp)

		pod2, err := clientset.CoreV1().Pods(namespace).Get("notapod", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", "notapod", namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				podName, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s created at %v\n", podName, namespace, pod2.CreationTimestamp)
		}
		fmt.Println()

		fmt.Println("------------------- Nodes -------------------")
		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		fmt.Printf("There are %d node in the cluster\n", len(nodes.Items))
		fmt.Println()
		fmt.Println("******************* List All *******************")
		fmt.Println()

		for _, node := range nodes.Items {
			fmt.Printf("| %6s |\n", node.GetName())
		}

		time.Sleep(2000 * time.Second)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
