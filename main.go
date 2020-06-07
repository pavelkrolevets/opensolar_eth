package main

import (
	"github.com/pavelkrolevets/opensolar_eth/db"
	"github.com/pavelkrolevets/opensolar_eth/handlers"
	"log"
	"net/http"
)

func main(){
	var store db.Store
	store.Path = "bolt_main"
	http.HandleFunc("/auth", handlers.Auth)
	http.HandleFunc("/new_user", handlers.NewUser)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
