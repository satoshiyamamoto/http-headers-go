package main

import (
	"fmt"
	"net/http"
)

func HttpHeadersHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s=%s\n", key, value)
	}
}

func main() {
	http.HandleFunc("/", HttpHeadersHandler)
	http.ListenAndServe(":8080", nil)
}
