package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// Equivalent to "subscribe" half in pub-sub model.
	// Create requests to connect that dials to a specific server and connects.
	// Format for address is different.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// Close connetion at the end of program.
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	// Print returned bs if any.
	if bs != nil {
		fmt.Println(string(bs))
	}
}
