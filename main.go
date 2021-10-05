package main

import (
	"log"
	"net/http"

	glserver "github.com/lnikon/glfs-pkg/pkg/server"

	httptransport "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

const (
	port = "8090"
)

func main() {
	// Cretate services and respective handlers
	algorithmService := glserver.NewAlgorithmService()
	algorithmHandler := httptransport.NewServer(
		glserver.MakeAlgorithmEndpoint(algorithmService),
		glserver.DecodeAlgorithmRequest,
		glserver.EncodeResponse,
	)

	computationService, err := glserver.NewComputationService()
	if err != nil {
		log.Fatal("Unable to create computation service!")
		return
	}

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
