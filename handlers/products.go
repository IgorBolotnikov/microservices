package handlers

import (
	"log"
	"net/http"

	"microservices/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		p.getProducts(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	list := data.GetProducts()
	err := list.ToJSON(writer)
	if err != nil {
		http.Error(writer, "Unable to encode product list", http.StatusInternalServerError)
	}
}
