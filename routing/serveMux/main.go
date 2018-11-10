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

func main() {
	var d dog
	var c cat

	mux := http.NewServeMux()
	// trailing slash makes url string still coverd by logic
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)

}
