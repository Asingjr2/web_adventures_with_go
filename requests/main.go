// Training code : https://github.com/GoesToEleven/golang-web-dev/blob/master/017_understanding-net-http-package/03_Request/05_Host_ContentLength/main.go
package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type madeUp int

// Makes madeUp type a type handler since it shares method with interface
func (m madeUp) ServeHttp(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	// tpl.ExecuteTemplate(res, "first.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var m madeUp
	// Below function taking port/url and a handler to deal with requests
	http.ListenAndServe(":8080", m)
}
