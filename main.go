package main

import (
	"log"
	"net/http"

	"github.com/alagha-go/go-amazon/lib/handler"
	_ "github.com/alagha-go/go-amazon/lib/variables"
)

var (
	PORT = ":7000"
)


func main() {
	err := http.ListenAndServe(PORT, handler.ServeMux)
	HandlError(err)
}


// handle errors by pannic
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}