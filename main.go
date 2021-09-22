package main

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	glserver "github.com/lnikon/glfs/pkg/server"
	// kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	// stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	port = "8090"
)

func main() {
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

	http.Handle("/algorithm", algorithmHandler)
	http.Handle("/computation", getAllComputationsHandler)

	// hello
	http.ListenAndServe(":"+port, nil)
}
