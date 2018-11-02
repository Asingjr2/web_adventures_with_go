// Experimenting with strings function.
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Seperator was space
	ts1 := "first; second[third]fourth{fifth}six       ! seven"
	fmt.Println(strings.Fields(ts1)[0])
	fmt.Println(strings.Fields(ts1)[1])
	fmt.Println(strings.Fields(ts1)[2])
	fmt.Println(strings.Fields(ts1)[3])
}
