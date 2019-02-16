package main

import (
  "log"
  "net/http"
  "github.com/camjw/stringsvc/stringsvc"
  httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := stringsvc.StringService{}

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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
