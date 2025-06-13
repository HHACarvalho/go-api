package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "CREATE")
	})

	mux.HandleFunc("GET /user/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "READ")
	})

	mux.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "READ ONE = " + id)
	})

	mux.HandleFunc("PUT /user/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "UPDATE")
	})

	mux.HandleFunc("DELETE /user/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "DELETE")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}