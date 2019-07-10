package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create a mux to hold url and handlers
	mux := http.NewServeMux()
	log.Println("created mux")

	// register handlers using HandleFunc
	mux.HandleFunc("/", root)
	mux.HandleFunc("/foo", foo)
	mux.HandleFunc("/bar", bar)
	log.Println("registered handlers")

	// register mux with server and listen for requests
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// HandlerFunc to be registered
func root(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "this is root")
}

func foo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "this is foo")
}

func bar(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "this is bar")
}
