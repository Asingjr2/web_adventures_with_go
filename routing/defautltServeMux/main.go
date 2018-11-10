package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

type cat int

func (d cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat")
}

// Func to be passed in handlefunc that does not require a handler type
func forHandleFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat")
}

func main() {
	var d dog
	var c cat
	http.Handle("/dog", d)
	http.Handle("/cat", c)

	// Does not require handle function, as it uses defaultMux automatically
	http.HandleFunc("/cat", forHandleFunc)

	// When nil is passed into this function defaultServeMux is used
	http.ListenAndServe(":8080", nil)
}
