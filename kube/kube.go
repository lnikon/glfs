package kube

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	upcxxv1alpha1types "github.com/lnikon/upcxx-operator/api/v1alpha1"
	upcxxv1alpha1clientset "github.com/lnikon/upcxx-operator/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/homedir"
	ctrl "sigs.k8s.io/controller-runtime"
)

func init() {
	if flag.Lookup("kubeconfig") == nil {
		if home := homedir.HomeDir(); home != "" {
			flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
		} else {
			flag.String("kubeconfig", "", "")
		}
	}
}

func createUpcxxClient() upcxxv1alpha1clientset.UPCXXInterface {
	flag.Parse()

	//kubeconfig := flag.Lookup("kubeconfig")
	//config, err := clientcmd.BuildConfigFromFlags("", kubeconfig.Value.String())
	config := ctrl.GetConfigOrDie()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	clientset, err := upcxxv1alpha1clientset.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return clientset.UPCXX("default")
}

// TODO: Review
//func GetPodsCount() int {
//	var kubeconfig *string
//	if home := homedir.HomeDir(); home != "" {
//		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
//	} else {
//		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig file")
//	}
//	flag.Parse()
//
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	return len(pods.Items)
//}

type UPCXXRequest struct {
	Name     string
	Replicas int32
}

type UPCXXResponse struct {
	Name     string
	Replicas int32
	IP       net.IP
}

func CreateUPCXX(req UPCXXRequest) error {
	upcxxClient := createUpcxxClient()

	groupVersionKind := schema.GroupVersionKind{}
	groupVersionKind.Group = upcxxv1alpha1types.GroupVersion.Group
	groupVersionKind.Version = upcxxv1alpha1types.GroupVersion.Version
	groupVersionKind.Kind = "UPCXX"

	apiVersion, kind := groupVersionKind.ToAPIVersionAndKind()

	upcxx := &upcxxv1alpha1types.UPCXX{
		TypeMeta: metav1.TypeMeta{
			Kind:       kind,
			APIVersion: apiVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: "default",
		},
		Spec: upcxxv1alpha1types.UPCXXSpec{
			StatefulSetName: req.Name,
			WorkerCount:     req.Replicas,
		},
		Status: upcxxv1alpha1types.UPCXXStatus{},
	}

	upcxx, err := upcxxClient.Create(upcxx)
	return err
}

func GetDeployment(name string) *UPCXXResponse {
	result := &UPCXXResponse{}

	upcxxClient := createUpcxxClient()
	deployement, err := upcxxClient.Get(name, metav1.GetOptions{})
	if err != nil {
		return nil
	}

	launcherSvc, err := upcxxClient.GetLauncherService(name)
	if err != nil {
		log.Printf("kube.GetDeployment failed to get launcher service\n")
		return nil
	}

	log.Printf("GetDeployment: Got launcher service: %v\n", launcherSvc)

	ip := net.IP{}
	if len(launcherSvc.Status.LoadBalancer.Ingress) > 0 {
		ip = net.ParseIP(launcherSvc.Status.LoadBalancer.Ingress[0].IP)
	} else {
		log.Println("GetDeployment: Empty ExternalIPs")
	}

	result = &UPCXXResponse{Name: deployement.Name, Replicas: deployement.Spec.WorkerCount, IP: ip}

	return result
}

func GetAllDeployments() []UPCXXResponse {
	result := []UPCXXResponse{}

	upcxxClient := createUpcxxClient()
	deploymentList, err := upcxxClient.List(metav1.ListOptions{})
	if err != nil {
		return result
	}

	for _, upcxx := range deploymentList.Items {
		response := UPCXXResponse{Name: upcxx.Spec.StatefulSetName, Replicas: upcxx.Spec.WorkerCount}

		launcherSvc, err := upcxxClient.GetLauncherService(upcxx.Spec.StatefulSetName)
		if err != nil {
			return []UPCXXResponse{}
		}

		if len(launcherSvc.Status.LoadBalancer.Ingress) > 0 {
			response.IP = net.ParseIP(launcherSvc.Status.LoadBalancer.Ingress[0].IP)
		} else {
			log.Println("GetDeployment: Empty ExternalIPs")
		}

		result = append(result, response)
	}

	return result
}

func DeleteDeployment(name string) error {
	upcxxClient := createUpcxxClient()
	_, err := upcxxClient.Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Error=%s\n", err.Error())
	}
	return err
}
