package main

import (
	"net/http"
	"os"

	log "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	glserver "github.com/lnikon/glfs/cmd/server"
	// kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	// stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	port = "8090"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// fieldKeys := []string{"method", "error"}
	// requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "string_service",
	// 	Name:      "request_count",
	// 	Help:      "Number of requests received",
	// }, fieldKeys)
	// requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "string_service",
	// 	Name:      "request_latency_microseconds",
	// 	Help:      "Total duration of requests in microseconds.",
	// }, fieldKeys)
	// countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "string_service",
	// 	Name:      "count_result",
	// 	Help:      "The result of each count method.",
	// }, []string{}) // no fields here

	var algorithmService glserver.AlgorithmService
	algorithmHandler := httptransport.NewServer(
		glserver.makeAlgorithmEndpoint(algorithmService),
		glserver.decodeAlgorithmRequest,
		glserver.encodeResponse,
	)

	var computationService glserver.ComputationService
	getAllComputationsHandler := httptransport.NewServer(
		glserver.makeGetAllComputationsEndpoint(computationService),
		glserver.decodeGetAllComputationsRequest,
		glserver.encodeResponse,
	)

	http.Handle("/algorithm", algorithmHandler)
	http.Handle("/computation", getAllComputationsHandler)

	logger.Log("err", http.ListenAndServe(":"+port, nil))
}
