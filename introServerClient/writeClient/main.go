package main

import (
	"fmt"
	"net"
)

func main() {
	// Create connection that dials to a specific server and connects
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// Close connetion at the end of program.
	defer conn.Close()

	// Write/ Print something to server prompt.
	fmt.Fprintln(conn, "I dialed you.") // Server set up with scanner function
}
