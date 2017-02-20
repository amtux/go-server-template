package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

func main() {
	log.Info("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
