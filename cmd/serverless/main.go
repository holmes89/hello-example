package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/holmes89/hello-example/hello"
	"go.opencensus.io/trace"
	"gocloud.dev/server"
	"gocloud.dev/server/sdserver"
)

func main() {
	addr := ":3000"
	options := &server.Options{
		RequestLogger: sdserver.NewRequestLogger(),
		// In production you will likely want to use trace.ProbabilitySampler
		// instead, since AlwaysSample will start and export a trace for every
		// request - this may be prohibitively slow with significant traffic.
		DefaultSamplingPolicy: trace.AlwaysSample(),
		Driver:                &server.DefaultDriver{},
	}

	mux := http.NewServeMux()
	languageService := hello.NewLanguageService()
	helloHandler := hello.NewHelloHandler(languageService)
	languageHandler := hello.NewLanguageHandler(languageService)

	mux.HandleFunc("/hello", helloHandler.Find)
	mux.HandleFunc("/languages", languageHandler.List)

	s := server.New(mux, options)
	fmt.Printf("Listening on %s\n", addr)

	err := s.ListenAndServe(addr)
	if err != nil {
		log.Fatal(err)
	}
}
