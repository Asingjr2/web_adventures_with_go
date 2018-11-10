// FuncMap is used to pass functiosn in to template when parsed
package main

import (
	"fmt"
	"os"
	"text/template"
)

type SomeNames struct {
	Mine  string
	Yours string
}

// Special map creation
var fm = template.FuncMap{
	"plusBowser": plusBow,
	"plusMario":  plusMar,
}

func plusBow(s string) string {
	s = s + "andBowser"
	return s
}

func plusMar(s string) string {
	s = s + "andMar"
	return s
}

func main() {
	myName := "Samus"
	yourName := "Random"
	// Creating instance of new type
	twoNames := SomeNames{myName, yourName}

	// Creating blank template..attaching functions...parsing files into contain template
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("someTemplate.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "someTemplate.gohtml", twoNames)
	if err != nil {
		fmt.Println(err)
	}
}
