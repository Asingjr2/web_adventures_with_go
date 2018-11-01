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

	scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)	splits characters

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
}
