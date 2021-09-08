package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var service = flag.String("s", "golang", "服务名")
var namespace = flag.String("n", "local", "命名空间")

func main() {
	flag.Parse()

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get ipaddress
	eps, err := clientset.CoreV1().Endpoints(*namespace).Get(context.TODO(), *service, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range eps.Subsets {
		for _, ip := range item.Addresses {
			fmt.Println(ip.IP)
		}
	}

}
