package server

import (
	"fmt"

	glkube "github.com/lnikon/glfs/kube"
)

type ComputationAllocationDescription struct {
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
}

type ComputationAllocation struct {
	Allocations []ComputationAllocationDescription
}

type ComputationAllocationServiceIfc interface {
	GetComputation(name string) (ComputationAllocationDescription, error)
	GetAllComputations() []ComputationAllocationDescription
	PostComputation(ComputationAllocationDescription) error
	DeleteComputation(name string) error
}

type ComputationAllocationService struct {
	computations ComputationAllocation
}

func NewComputationService() (ComputationAllocationServiceIfc, error) {
	computationService := &ComputationAllocationService{}
	return computationService, nil
}

func (c *ComputationAllocationService) GetAllComputations() []ComputationAllocationDescription {
	upcxxList := glkube.GetAllDeployments()
	descriptions := []ComputationAllocationDescription{}
	for _, upcxx := range upcxxList.Items {
		descriptions = append(descriptions, ComputationAllocationDescription{
			Name:     upcxx.Spec.StatefulSetName,
			Replicas: upcxx.Spec.WorkerCount,
		})
	}

	return descriptions
}

func (c *ComputationAllocationService) GetComputation(name string) (ComputationAllocationDescription, error) {
	upcxx := glkube.GetDeployment(name)
	if upcxx == nil {
		return ComputationAllocationDescription{}, fmt.Errorf("resource does not exists")
	}

	return ComputationAllocationDescription{
		Name:     upcxx.Spec.StatefulSetName,
		Replicas: upcxx.Spec.WorkerCount,
	}, nil
}

func (c *ComputationAllocationService) PostComputation(description ComputationAllocationDescription) error {
	upcxxReq := glkube.UPCXXRequest{Name: description.Name, Replicas: description.Replicas}
	if err := glkube.CreateUPCXX(upcxxReq); err != nil {
		return err
	}

	c.computations.Allocations = append(c.computations.Allocations, description)
	return nil
}

func (c *ComputationAllocationService) DeleteComputation(name string) error {
	return glkube.DeleteDeployment(name)
}
