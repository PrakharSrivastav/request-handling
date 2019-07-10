package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandleFunc(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "this is serving from http-HandleFunc")
}
