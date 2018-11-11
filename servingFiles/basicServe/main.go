package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// io.WriteString(w, `
	// <! -- serving from site -->
	// <img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	// `)

	// Setting type to json which can be read differnetly.  Would return error as html
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `
	{"Name":"Smaus", "Type":"Hero"}
	`)
}

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}
