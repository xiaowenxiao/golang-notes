package main

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// 通过集群内部配置创建 k8s 配置信息，通过 KUBERNETES_SERVICE_HOST 和 KUBERNETES_SERVICE_PORT 环境变量方式获取
	// 若集群使用 TLS 认证方式，则默认读取集群内部 tokenFile 和 CAFile
	// tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	// rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// 根据指定的 config 创建一个新的 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string) 返回 PodInterface
		// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
		// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pod 列表
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the k8s cluster\n", len(pods.Items))

		// 获取指定 namespace 中的 Pod 列表信息
		namespce := "local"
		pods, err = clientset.CoreV1().Pods(namespce).List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nThere are %d pods in namespaces %s\n", len(pods.Items), namespce)
		for _, pod := range pods.Items {
			fmt.Printf("Name: %s, Status: %s, CreateTime: %s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.CreationTimestamp)
		}

		// 获取所有的 Namespaces 列表信息
		ns, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		nss := ns.Items
		fmt.Printf("\nThere are %d namespaces in cluster\n", len(nss))
		for _, ns := range nss {
			fmt.Printf("Name: %s, Status: %s, CreateTime: %s\n", ns.ObjectMeta.Name, ns.Status.Phase, ns.CreationTimestamp)
		}

		time.Sleep(10 * time.Second)
	}
}
