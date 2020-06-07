package main

import (
	"github.com/boltdb/bolt"
	"log"
	"net/http"
	"github.com/pavelkrolevets/opensolar_eth/handlers"
)

func main(){

	http.HandleFunc("/auth", handlers.Auth)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
