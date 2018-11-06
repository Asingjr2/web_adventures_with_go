package main

import "fmt"

func main() {
	name := "Samus"

	tpl := `
	<!DOCTYPE html>
	<head>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
