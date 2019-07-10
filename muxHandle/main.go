package main

import (
	"fmt"
	"log"
	"net/http"
)

// myHandler implements ServeHTTP, so it is valid
type myHandler struct{}
func (m *myHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "serving via mux-Handle")
}

func main() {
	// handler
	h := new(myHandler)

	// create mux and register handler
	mux := http.NewServeMux()
	mux.Handle("/", h)

	// register mux with server and listen for requests
	log.Fatal(http.ListenAndServe(":8080", mux))
}
