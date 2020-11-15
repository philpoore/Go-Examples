package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world")
	fmt.Print(".")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, value := range headers {
			fmt.Printf("%s = %s\n", name, value)
		}
	}
}

func main() {
	port := ":20000"
	fmt.Println("Starting server")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(port, nil)
}
