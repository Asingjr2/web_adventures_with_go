package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Parsefiles returns a type template (e.g. container for templates)
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Writing information in tpl to screen
	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating new file", err)
	}
	defer nf.Close()

	// Executes takes a writer and data and writes infor,ation
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
