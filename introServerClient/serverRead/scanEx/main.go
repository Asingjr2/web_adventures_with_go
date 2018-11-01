package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	firstString := "I got a feeling\n...hooked on believing...\nBrandiiiiiii"

	// Takes in a io.Reader type or interface
	scanner := bufio.NewScanner(strings.NewReader(firstString))

	count := 0
	for scanner.Scan() {
		fmt.Println(count)
		ln := scanner.Text()
		fmt.Println(ln)
		// Including below line would advance scanner forward skipping middle part
		// fmt.Println(scanner.Scan())
		count++
	}
}
