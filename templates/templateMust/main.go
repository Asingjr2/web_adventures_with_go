package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// Must does error templating and takes what glob returns
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)

	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
