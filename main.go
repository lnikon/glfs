package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/log"
	mux "github.com/gorilla/mux"
	glserver "github.com/lnikon/glfs-pkg/pkg/server"
)

const (
	hostname = ":"
	port     = "8090"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	// Cretate services and respective handlers
	algorithmService := glserver.NewAlgorithmService()
	algorithmHandler := httptransport.NewServer(
		glserver.MakeAlgorithmEndpoint(algorithmService),
		glserver.DecodeAlgorithmRequest,
		glserver.EncodeResponse,
	)

	computationService, err := glserver.NewComputationService()
	computationService = glserver.LoggingMiddleware{Next: computationService, Logger: logger}
	if err != nil {
		logger.Log("unable to create computation service")
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
	logger.Log("host", hostname, "port", port)
	http.ListenAndServe(hostname+port, router)
}
