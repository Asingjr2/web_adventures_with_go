// Creating TCP server adn writing into connection which we can see through telnet
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Creates server, type in tcp here on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// Defer the close of tcp server so it does not stay open
	defer li.Close()

	for {
		// Accept waits for connection request and then accepts connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Write method below puts information into through the connection
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		fmt.Printf("%T", li) // returns type net listener

		conn.Close() // Ensuring close of connection
	}
}
