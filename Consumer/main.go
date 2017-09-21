package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("DRP_CF_HTTP_PORT")
	url := os.Getenv("DRP_CF_HTTP_ADDR")
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
