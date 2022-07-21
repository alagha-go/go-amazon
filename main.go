package main

import (
	"log"
	"net/http"

	_ "github.com/alagha-go/go-amazon/lib/variables"
	"github.com/alagha-go/go-amazon/lib/handler"
)

var (
	PORT = ":8000"
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