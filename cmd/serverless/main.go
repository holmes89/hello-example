package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/holmes89/hello-example/hello"
	"gocloud.dev/server"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()
	languageService := hello.NewLanguageService()
	helloHandler := hello.NewHelloHandler(languageService)
	languageHandler := hello.NewLanguageHandler(languageService)

	mux.HandleFunc("/hello", helloHandler.Find)
	mux.HandleFunc("/languages", languageHandler.List)
	mux.HandleFunc("/", debugHandler)

	s := server.New(mux, nil)
	fmt.Printf("Listening on %s\n", addr)

	err := s.ListenAndServe(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %+v\n", r.URL)
	fmt.Fprint(w, r.URL)
}
