package main

import (
	"encoding/json"
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
	r.HandleFunc("/user", User).Methods("Get")
	r.HandleFunc("/.well-known/live", Live).Methods("Get")
	r.HandleFunc("/.well-known/ready", Ready).Methods("Get")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", url, port), r))

}

// User func
func User(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["id"] = 1
	response["name"] = "john"
	val, err := json.Marshal(&response)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(val)
}

// Live func
func Live(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(204)
}

// Ready func
func Ready(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(204)
}
