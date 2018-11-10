package main

import (
	"io"
	"net/http"
)

type madeUp int

func (m madeUp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/dog":
		io.WriteString(res, "was a dog")
	case "/cat":
		io.WriteString(res, "was a cat")
	}
}

func main() {
	var mu madeUp
	http.ListenAndServe(":8080", mu)
}
