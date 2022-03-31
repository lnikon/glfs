package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type GetComputationRequest struct {
	Name string `json:"name"`
}

type GetComputationResponse struct {
	Computation ComputationAllocationDescription `json:"computation"`
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
}

func MakeGetComputationEndpoint(svc ComputationAllocationServiceIfc) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetComputationRequest)
		computation, err := svc.GetComputation(req.Name)
		if err != nil {
			return nil, err
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
		err := svc.PostComputation(ComputationAllocationDescription{Name: req.Name, Replicas: req.Replicas})
		if err != nil {
			return nil, err
		}

		return PostComputationResponse{}, nil
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
