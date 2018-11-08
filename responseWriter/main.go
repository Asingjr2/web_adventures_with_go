package main

import (
	"fmt"
	"net/http"
)

// Which will be used as interface with below method
type madeUp int

func (m madeUp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Sets header object/type
	rw.Header().Set("Arthur Key", "some info")
	rw.Header().Set("Content-Type", "text/text; charset=utf=8")
	fmt.Fprintln(rw, "<h1>I AM AMAZING</h1>")
}

func main() {
	var mu madeUp
	http.ListenAndServe(":8080", mu)
}
