package main

import (
	"net/http"

	glserver "github.com/lnikon/glfs-pkg/server"

	httptransport "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

const (
	port = "8090"
)

func main() {
	// Cretate services and respective handlers
	var algorithmService glserver.AlgorithmService
	algorithmHandler := httptransport.NewServer(
		glserver.MakeAlgorithmEndpoint(algorithmService),
		glserver.DecodeAlgorithmRequest,
		glserver.EncodeResponse,
	)

	var computationService glserver.ComputationService
	getAllComputationsHandler := httptransport.NewServer(
		glserver.MakeGetAllComputationsEndpoint(computationService),
		glserver.DecodeGetAllComputationsRequest,
		glserver.EncodeResponse,
	)

	postComputationHandler := httptransport.NewServer(
		glserver.MakePostComputationEndpoint(computationService),
		glserver.DecodePostComputationRequest,
		glserver.EncodeResponse,
	)

	// Do routing staff
	router := mux.NewRouter()
	router.Methods("GET").Path("/algorithm").Handler(algorithmHandler)
	router.Methods("GET").Path("/computation").Handler(getAllComputationsHandler)
	router.Methods("POST").Path("/computation").Handler(postComputationHandler)

	// Start to listen for incoming requests
	http.ListenAndServe(":"+port, router)
}
