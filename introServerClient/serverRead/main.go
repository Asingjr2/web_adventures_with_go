// Bufion support I/O functions with buffers for connections.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Creating server type tcp on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// Defering close of server operation
	defer li.Close()

	for {
		// Listening for and creating a connection
		// Creating infinite loop so connection has to be closed manually
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// Running directions to handle inforamtion coming into connection
		go handle(conn)
	}
}

// Below initiates infinite loop to print information while in coming through channel
func handle(conn net.Conn) {
	// Scanner takes in a type io.Reader which conn is (as well as writer)
	// Returns scanner object
	scanner := bufio.NewScanner(conn)
	// Scan() advances to next token to see if exits.  returns false if not
	// Scanner.scan returns a bool...so for true (or while true)
	for scanner.Scan() {
		// Text returns information received by scan
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("Code got here!!!")
}
