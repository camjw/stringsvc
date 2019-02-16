package main

import (
	"github.com/camjw/stringsvc/stringsvc"
	"net/http"
	"os"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var svc stringsvc.Service
	svc = stringsvc.StringService{}
	svc = stringsvc.LoggingMiddleware{logger, svc}
  svc = stringsvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	port := ":" + os.Getenv("TARGET_PORT")
	if port == ":" {
		port = ":8080"
	}

	uppercaseHandler := httptransport.NewServer(
		stringsvc.MakeUppercaseEndpoint(svc),
		stringsvc.DecodeUppercaseRequest,
		stringsvc.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		stringsvc.MakeCountEndpoint(svc),
		stringsvc.DecodeCountRequest,
		stringsvc.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, nil))
}
