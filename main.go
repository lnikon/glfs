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

	getComputationHandler := httptransport.NewServer(
		glserver.MakeGetComputationEndpoint(computationService),
		glserver.DecodeGetComputationRequest,
		glserver.EncodeResponse,
	)

	postComputationHandler := httptransport.NewServer(
		glserver.MakePostComputationEndpoint(computationService),
		glserver.DecodePostComputationRequest,
		glserver.EncodeResponse,
	)

	// Do routing staff
	router := mux.NewRouter()

	// Endpoint for algorithms
	router.Methods("GET").Path("/algorithm").Handler(algorithmHandler)

	// Endpoints for computations
	// GET endpoints
	computationRouter := router.PathPrefix("/computation").Subrouter()
	computationRouter.Methods("GET").Path("/{name}").Handler(getComputationHandler)
	computationRouter.Methods("GET").Path("/").Handler(getAllComputationsHandler)

	// POST endpoints
	computationRouter.Methods("POST").Path("/").Handler(postComputationHandler)

	// Start to listen for incoming requests
	log.Default().Printf("Listening on %s:%s\n", ":", port)
	http.ListenAndServe(":"+port, router)
}
