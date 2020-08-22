package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Gocker!")
}

func main() {
	fmt.Println("Gocker Version 1.0")

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":9090", nil)
}
