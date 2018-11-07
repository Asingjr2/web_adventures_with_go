// Copied from training understarnding-net-http
// Inforamtion is sent in this and added in other code
package main

import (
	"html/template"
	"log"
	"net/http" // contains handler, req types and methods
	"net/url"  // contains url type
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Must be run in order to get some field information
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// Defining literal struct with creation
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form, // correlates to submissions
		req.Header,
		req.Host,
		req.ContentLength,
	}
	// Function writes object into a template.  Below writes data information into html template
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	// d handles the request in logic created above
	http.ListenAndServe(":8080", d)
}
