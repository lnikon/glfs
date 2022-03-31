package server

import (
	"fmt"
	"net"

	glkube "github.com/lnikon/glfs/kube"
)

type ComputationAllocationDescription struct {
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
	IP       net.IP `json:"ip"`
}

type ComputationAllocation struct {
	Allocations []ComputationAllocationDescription `json:"allocations"`
}

type ComputationAllocationServiceIfc interface {
	GetComputation(name string) (ComputationAllocationDescription, error)
	GetAllComputations() []ComputationAllocationDescription
	PostComputation(ComputationAllocationDescription) error
	DeleteComputation(name string) error
}

type ComputationAllocationService struct {
	computations ComputationAllocation `json:"computations"`
}

func NewComputationService() (ComputationAllocationServiceIfc, error) {
	computationService := &ComputationAllocationService{}
	return computationService, nil
}

func (c *ComputationAllocationService) GetAllComputations() []ComputationAllocationDescription {
	upcxxResponse := glkube.GetAllDeployments()
	descriptions := []ComputationAllocationDescription{}
	for _, upcxx := range upcxxResponse {
		descriptions = append(descriptions, ComputationAllocationDescription{
			Name:     upcxx.Name,
			Replicas: upcxx.Replicas,
			IP:       upcxx.IP,
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
		Name:     upcxx.Name,
		Replicas: upcxx.Replicas,
		IP:       upcxx.IP,
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
