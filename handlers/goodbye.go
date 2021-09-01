package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (goodbye *Goodbye) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	goodbye.log("Goodbye World!")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(rw, "There was an error!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Goodbye, %s", data)
}

func (goodbye *Goodbye) log(text string) {
	goodbye.logger.Println(text)
}