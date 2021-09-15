package main

type Computation struct {
	algorithm Algorithm
}

type ComputationService struct {
}

func (c *ComputationService) GetAllComputations() []Computation {
	return []Computation{{algorithm: Kruskal}}
}
