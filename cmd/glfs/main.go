package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/log"
	mux "github.com/gorilla/mux"
	glserver "github.com/lnikon/glfs/server"
)

const (
	hostname = ":"
	port     = "8090"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	computationService, err := glserver.NewComputationService()
	computationService = glserver.LoggingMiddleware{Next: computationService, Logger: logger}
	if err != nil {
		logger.Log("unable to create computation service")
		return
	}

	getAllocationHandler := httptransport.NewServer(
		glserver.MakeGetAllocationEndpoint(computationService),
		glserver.DecodeGetAllocationRequest,
		glserver.EncodeResponse,
	)

	getAllAllocationsHandler := httptransport.NewServer(
		glserver.MakeGetAllAllocationsEndpoint(computationService),
		glserver.DecodeGetAllAllocationsRequest,
		glserver.EncodeResponse,
	)

	postAllocationHandler := httptransport.NewServer(
		glserver.MakePostAllocationEndpoint(computationService),
		glserver.DecodePostAllocationRequest,
		glserver.EncodeResponse,
	)

	deleteAllocationHandler := httptransport.NewServer(
		glserver.MakeDeleteAllocationEndpoint(computationService),
		glserver.DecodeDeleteAllocationRequest,
		glserver.EncodeResponse,
	)

	router := mux.NewRouter()
	allocationRouter := router.PathPrefix("/allocation").Subrouter()
	allocationRouter.Methods("GET").Path("/result").Handler(getAllocationHandler)
	allocationRouter.Methods("GET").Path("/result/all").Handler(getAllAllocationsHandler)
	allocationRouter.Methods("POST").Handler(postAllocationHandler)
	allocationRouter.Methods("DELETE").Handler(deleteAllocationHandler)

	// Start to listen for incoming requests
	// VAGAGTODO: Update logging
	logger.Log("host", hostname, "port", port)
	err = http.ListenAndServe(hostname+port, router)
	if err != nil {
		logger.Log("Error", err)
	}
}
