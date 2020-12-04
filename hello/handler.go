package hello

import (
	"encoding/json"
	"net/http"
)

type HelloHandler struct {
	service *LanguageService
}

func NewHelloHandler(service *LanguageService) *HelloHandler {
	return &HelloHandler{
		service: service,
	}
}

func (h *HelloHandler) Find(w http.ResponseWriter, r *http.Request) {
	languages := r.URL.Query()["language"]
	res := h.service.GetHellos(languages)
	encodeResponse(w, res)
	return
}

type LanguageHandler struct {
	service *LanguageService
}

func NewLanguageHandler(service *LanguageService) *LanguageHandler {
	return &LanguageHandler{
		service: service,
	}
}

func (h *LanguageHandler) List(w http.ResponseWriter, r *http.Request) {
	res := h.service.ListLanguages()
	encodeResponse(w, res)
	return
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := enc.Encode(response); err != nil {
		panic("unable to encode response")
	}
}
