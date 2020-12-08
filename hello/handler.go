package hello

import (
	"encoding/json"
	"net/http"
)

// WordHandler holds service for finding language results
type WordHandler struct {
	service *LanguageService
}

// NewHelloHandler instanciates a new handler
func NewHelloHandler(service *LanguageService) *WordHandler {
	return &WordHandler{
		service: service,
	}
}

// Find will fetch hello in a given list of languages
func (h *WordHandler) Find(w http.ResponseWriter, r *http.Request) {
	languages := r.URL.Query()["language"]
	res := h.service.GetHellos(languages)
	encodeResponse(w, res)
}

// LanguageHandler will list all languages in the system
type LanguageHandler struct {
	service *LanguageService
}

// NewLanguageHandler instanciates a new handler
func NewLanguageHandler(service *LanguageService) *LanguageHandler {
	return &LanguageHandler{
		service: service,
	}
}

// List will return all languages
func (h *LanguageHandler) List(w http.ResponseWriter, r *http.Request) {
	res := h.service.ListLanguages()
	encodeResponse(w, res)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := enc.Encode(response); err != nil {
		panic("unable to encode response")
	}
}
