package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/holmes89/hello-example/hello"
)

func main() {
	addr := ":3000"

	mux := http.NewServeMux()
	languageService := hello.NewLanguageService()
	helloHandler := hello.NewHelloHandler(languageService)
	languageHandler := hello.NewLanguageHandler(languageService)

	mux.HandleFunc("/hello", helloHandler.Find)
	mux.HandleFunc("/languages", languageHandler.List)

	fmt.Printf("listening on %s\n", addr)

	log.Fatal(gateway.ListenAndServe(addr, mux))
}
