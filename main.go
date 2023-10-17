package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// Get the IP address of the server's connected user.
func getUserIP(httpWriter http.ResponseWriter, httpServer *http.Request) (net.IP, error) {
	var userIP string
	if len(httpServer.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = httpServer.Header.Get("CF-Connecting-IP")
		return net.ParseIP(userIP), nil
	} else if len(httpServer.Header.Get("X-Forwarded-For")) > 1 {
		userIP = httpServer.Header.Get("X-Forwarded-For")
		return net.ParseIP(userIP), nil
	} else if len(httpServer.Header.Get("X-Real-IP")) > 1 {
		userIP = httpServer.Header.Get("X-Real-IP")
		return net.ParseIP(userIP), nil
	} else {
		userIP = httpServer.RemoteAddr
		if strings.Contains(userIP, ":") {
			return net.ParseIP(strings.Split(userIP, ":")[0]), nil
		} else {
			return net.ParseIP(userIP), nil
		}
	}
	return nil, fmt.Errorf("No valid ip found")
}

func writeHeaders(w io.Writer, req *http.Request, realIp net.IP) {

	fmt.Fprintf(w, "Hview\n\n\n")

	fmt.Fprintf(w, "Time:\n-------------------\n")
	fmt.Fprintf(w, "%s\n\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(w, "Header:\n-------------------\n")

	keys := make([]string, 0, len(req.Header))
	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintf(w, "%s: %s\n", k, req.Header[k])
	}

	fmt.Fprintf(w, "\nRemode ADDR:\n-------------------\n")
	fmt.Fprintf(w, "%s\n", req.RemoteAddr)

	fmt.Fprintf(w, "\nReal-USER-IP:\n-------------------\n")
	fmt.Fprintf(w, "%s\n", realIp)
	//
	fmt.Fprintf(w, "\nMethod:\n-------------------\n")
	fmt.Fprintf(w, "%s\n", req.Method)
}

func headers(w http.ResponseWriter, req *http.Request) {
	// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.
	ip, err := getUserIP(w, req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	writeHeaders(w, req, ip)
	writeHeaders(os.Stdout, req, ip)

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
