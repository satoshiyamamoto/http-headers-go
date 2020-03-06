package main

import (
	"fmt"
	"net/http"
	"sort"
)

func HttpHeadersHandler(w http.ResponseWriter, r *http.Request) {
	names := make([]string, 0, len(r.Header))
	for n := range r.Header {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Fprintf(w, "%s=%s\n", name, r.Header.Get(name))
	}
}

func main() {
	http.HandleFunc("/", HttpHeadersHandler)
	http.ListenAndServe(":8080", nil)
}
