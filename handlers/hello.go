package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.log("Hello World!")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "There was an error!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello, %s", data)
}

func (hello *Hello) log(text string) {
	hello.logger.Println(text)
}