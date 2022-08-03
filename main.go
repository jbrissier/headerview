package main

import (
	"fmt"
	"net/http"
	"os"
)

func headers(w http.ResponseWriter, req *http.Request) {
	// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	port := os.Getenv("H_VIEW_PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", headers)
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
