package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/log"
	mux "github.com/gorilla/mux" glserver "github.com/lnikon/glfs/server"
)

const (
	hostname = ":"
	port     = "8080"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

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
	computationRouter := router.PathPrefix("/computation").Subrouter()
	computationRouter.Methods("GET").Path("/name").Handler(getComputationHandler)
	computationRouter.Methods("GET").Path("/").Handler(getAllComputationsHandler)
	computationRouter.Methods("POST").Handler(postComputationHandler)

	// Start to listen for incoming requests
	// VAGAGTODO: Update logging
	logger.Log("host", hostname, "port", port)
	http.ListenAndServe(hostname+port, router)
}
