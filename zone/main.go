package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/zone", func(w http.ResponseWriter, r *http.Request) {})
	fmt.Println("starting http server on :4000")
	
	http.ListenAndServe(":4000", r)
}
