package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Glob takes a bunch of files
	tpl, err := template.ParseGlob("templates/*")

	err = tpl.Execute(os.Stdout, nil)

	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
