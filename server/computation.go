package server

import (
	"fmt"
	"net"

	glkube "github.com/lnikon/glfs/kube"
)

type AllocationDescription struct {
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`

	// Set in a result of GET request
	IP net.IP `json:"ip"`
}

type ComputationAllocation struct {
	Allocations []AllocationDescription `json:"allocations"`
}

type ComputationAllocationServiceIfc interface {
	GetAllocation(name string) (AllocationDescription, error)
	GetAllAllocations() []AllocationDescription
	PostAllocation(AllocationDescription) error
	DeleteAllocation(name string) error
}

type ComputationAllocationService struct {
	allocations ComputationAllocation
}

func NewComputationService() (ComputationAllocationServiceIfc, error) {
	computationService := &ComputationAllocationService{}
	return computationService, nil
}

func (c *ComputationAllocationService) GetAllocation(name string) (AllocationDescription, error) {
	upcxx := glkube.GetDeployment(name)
	if upcxx == nil {
		return AllocationDescription{}, fmt.Errorf("resource does not exists")
	}

	return AllocationDescription{
		Name:     upcxx.Name,
		Replicas: upcxx.Replicas,
		IP:       upcxx.IP,
	}, nil
}

func (c *ComputationAllocationService) GetAllAllocations() []AllocationDescription {
	upcxxResponse := glkube.GetAllDeployments()
	descriptions := []AllocationDescription{}
	for _, upcxx := range upcxxResponse {
		descriptions = append(descriptions, AllocationDescription{
			Name:     upcxx.Name,
			Replicas: upcxx.Replicas,
			IP:       upcxx.IP,
		})
	}

	return descriptions
}

func (c *ComputationAllocationService) PostAllocation(description AllocationDescription) error {
	upcxxReq := glkube.UPCXXRequest{Name: description.Name, Replicas: description.Replicas}
	if err := glkube.CreateUPCXX(upcxxReq); err != nil {
		return err
	}

	c.allocations.Allocations = append(c.allocations.Allocations, description)
	return nil
}

func (c *ComputationAllocationService) DeleteAllocation(name string) error {
	return glkube.DeleteDeployment(name)
}
