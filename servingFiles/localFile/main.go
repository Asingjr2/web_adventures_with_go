// Files serves local files into webpage
package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("sunset11.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}
