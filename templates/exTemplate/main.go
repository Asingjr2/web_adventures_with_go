package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("first.gohtml")

	err = tpl.Execute(os.Stdout, nil)

	tpl, err = tpl.ParseFiles("second.gohtml", "third.gohtml")
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "first.gohtml", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)

	// If multiple files are included ni tpl you can declare which to execute
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
