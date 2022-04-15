package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type GetAllocationRequest struct {
	Name string `json:"name"`
}

type GetAllocationResponse struct {
	Computation AllocationDescription `json:"computation"`
	Error       string                `json:"error,omitempty"`
}

type GetAllAllocationsRequest struct {
}

type GetAllAllocationsResponse struct {
	Computations []AllocationDescription `json:"computations"`
	Error        string                  `json:"error,omitempty"`
}

type PostAllocationRequest struct {
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
}

type PostAllocationResponse struct {
	Result string `json:"result,omitempty"`
}

type DeleteAllocationRequest struct {
	Name string `json:"name"`
}

type DeleteAllocationResponse struct {
	Error string `json:"error"`
}

func MakeGetAllocationEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllocationRequest)
		computation, err := svc.GetAllocation(req.Name)
		if err != nil {
			return GetAllocationResponse{Error: err.Error()}, nil
		}

		return GetAllocationResponse{Computation: computation}, nil
	}
}

func DecodeGetAllocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := GetAllocationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func MakeGetAllAllocationsEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return GetAllAllocationsResponse{Computations: svc.GetAllAllocations()}, nil
	}
}

func DecodeGetAllAllocationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllAllocationsRequest{}, nil
}

func MakePostAllocationEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PostAllocationRequest)
		if len(req.Name) == 0 || req.Replicas <= 0 {
			return nil, fmt.Errorf("Empty computation name=%s or wrong replicas=%d\n", req.Name, req.Replicas)
		}

		err := svc.PostAllocation(AllocationDescription{Name: req.Name, Replicas: req.Replicas})
		if err != nil {
			return PostAllocationResponse{Result: err.Error()}, nil
		}

		return PostAllocationResponse{Result: "Allocating"}, nil
	}
}

func DecodePostAllocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := PostAllocationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func MakeDeleteAllocationEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAllocationRequest)
		err := svc.DeleteAllocation(req.Name)
		if err != nil {
			return DeleteAllocationResponse{Error: err.Error()}, nil
		}

		return DeleteAllocationResponse{}, nil
	}
}

func DecodeDeleteAllocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := DeleteAllocationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
