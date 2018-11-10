package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("data.gohtml"))
}

func main() {
	// Pass only one piece of data
	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", 10000)
	if err != nil {
		log.Fatalln(err)
	}
}
