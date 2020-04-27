package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// use gofmt main.go        to format the code
// use goimports main.go    to import necessary files
// use go vet main.go       to check the code for errors

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Packt")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", exampleHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}
