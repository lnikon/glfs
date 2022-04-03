package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type GetComputationRequest struct {
	Name string `json:"name"`
}

type GetComputationResponse struct {
	Computation ComputationAllocationDescription `json:"computation"`
	Error       string                           `json:"error,omitempty"`
}

type GetAllComputationsRequest struct {
}

type GetAllComputationsResponse struct {
	Computations []ComputationAllocationDescription `json:"computations"`
}

type PostComputationRequest struct {
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
}

type PostComputationResponse struct {
	Result string `json:"result,omitempty"`
}

func MakeGetComputationEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetComputationRequest)
		computation, err := svc.GetComputation(req.Name)
		if err != nil {
			return GetComputationResponse{Error: err.Error()}, nil
		}

		return GetComputationResponse{Computation: computation}, nil
	}
}

func DecodeGetComputationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := GetComputationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func MakeGetAllComputationsEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return GetAllComputationsResponse{Computations: svc.GetAllComputations()}, nil
	}
}

func DecodeGetAllComputationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllComputationsRequest{}, nil
}

func MakePostComputationEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PostComputationRequest)
		if len(req.Name) == 0 || req.Replicas <= 0 {
			return nil, fmt.Errorf("Empty computation name=%s or wrong replicas=%d\n", req.Name, req.Replicas)
		}

		err := svc.PostComputation(ComputationAllocationDescription{Name: req.Name, Replicas: req.Replicas})
		if err != nil {
			return PostComputationResponse{Result: err.Error()}, nil
		}

		return PostComputationResponse{Result: "Allocating"}, nil
	}
}

func DecodePostComputationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := PostComputationRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
