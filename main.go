package main

import (
	"github.com/camjw/stringsvc/stringsvc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc stringsvc.Service
	svc = stringsvc.StringService{}
  svc = stringsvc.LoggingMiddleware{logger, svc}

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

	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, nil))
}
