package server

import (
	"fmt"
)

type Computation struct {
	algorithm Algorithm
}

type ComputationService struct {
}

func (c *ComputationService) GetAllComputations() []Computation {
	return []Computation{{algorithm: Kruskal}}
}

func (c *ComputationService) PostComputation(request *postComputationRequest) postComputationResponse {
	// glkube.CreateDeployment()
	fmt.Printf("Post computation request called for %s", request.Algorithm)
	return postComputationResponse{}
}
