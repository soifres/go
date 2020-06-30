package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	fmt.Fprint(r, m)
}

func main() {
	msgHandler := msg("Elya, Hello from web server.)))")
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", msgHandler)

	// http.HandlerFunc("/about", func (r http.ResponseWriter, w *http.Request) {

	// })

	// http.ListenAndServe("localhost:8080", nil)
}
