package main

import (
	"employe/pkg/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
