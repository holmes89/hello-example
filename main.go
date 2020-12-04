package main

import (
	"log"
	"net/http"

	"github.com/holmes89/hello-example/hello"
)

func main() {
	languageService := hello.NewLanguageService()
	helloHandler := hello.NewHelloHandler(languageService)
	languageHandler := hello.NewLanguageHandler(languageService)

	http.HandleFunc("/hello", helloHandler.Find)
	http.HandleFunc("/languages", languageHandler.List)
	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
