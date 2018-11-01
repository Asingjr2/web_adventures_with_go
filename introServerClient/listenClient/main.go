package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	// Print returned bs if any
	if bs != nil {
		fmt.Println(string(bs))
	}
}
