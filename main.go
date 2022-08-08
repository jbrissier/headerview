package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func headers(w http.ResponseWriter, req *http.Request) {
	// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.

	fmt.Fprintf(w, "Hview\n\n\n")

	fmt.Fprintf(w, "Time:\n-------------------\n")
	fmt.Fprintf(w, "%s\n\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "Header:\n-------------------\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "\nIP:\n-------------------\n")
	fmt.Fprintf(w, "%s\n", req.RemoteAddr)
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
