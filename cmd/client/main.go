package main

import (
	"net/http"
	"os"

	log "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	// kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	// stdprometheus "github.com/prometheus/client_golang/prometheus"
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

	var svc AlgorithmService
	algorithmHandler := httptransport.NewServer(
		makeAlgorithmEndpoint(svc),
		decodeAlgorithmRequest,
		encodeResponse,
	)

	http.Handle("/algorithm", algorithmHandler)
	logger.Log("msg", "HTTP", "addr", ":8090")
	logger.Log("err", http.ListenAndServe(":8091", nil))
}
