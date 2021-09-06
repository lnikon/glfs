package main 

import (
	"context"
	"net/http"
	"encoding/json"

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

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
