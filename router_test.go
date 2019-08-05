package gorouter

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewRouter(t *testing.T) {
	mux := New()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello,csuhan")
	})
	http.ListenAndServe(":8080", mux)
}
