package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
