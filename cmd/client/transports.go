package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type algorithmRequest struct {
}

type algorithmResponse struct {
	Algorithm []Algorithm `json:"algorithm"`
}

func makeAlgorithmEndpoint(svc AlgorithmService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		a := svc.Algorithm()
		return algorithmResponse{a}, nil
	}
}

func decodeAlgorithmRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return algorithmRequest{}, nil
}

type GetAllComputationsRequest struct {
}

type GetAllComputationsResponse struct {
	computations []Computation
}

func makeGetAllComputationsEndpoint(svc ComputationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		s := svc.GetAllComputations()
		return GetAllComputationsResponse{s}, nil
	}
}

func decodeGetAllComputationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllComputationsRequest{}, nil
}

// Universal encoder for all responses
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
