package main

import (
	"fmt"
	"net/http"

	"github.com/csuhan/gorouter"
)

func main() {
	router := gorouter.New()
	router.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("CSUHAN GET"))
	})
	router.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,%s", "csuhan")
	})
	http.ListenAndServe(":8080", router)
}
