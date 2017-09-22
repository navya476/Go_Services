package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var port string
	var url string
	if os.Getenv("DRP_CF_HTTP_ADDR") != "" {
		url = os.Getenv("DRP_CF_HTTP_ADDR")
	} else {
		url = "localhost"
	}
	if os.Getenv("DRP_CF_HTTP_PORT") != "" {
		port = os.Getenv("DRP_CF_HTTP_PORT")
	} else {
		port = "8085"
	}
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/live", Live).Methods("Get")
	r.HandleFunc("/.well-known/ready", Ready).Methods("Get")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", url, port), r))

}

// Live func
func Live(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(204)
}

// Ready func
func Ready(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(204)
}
