package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text() // reads information in connection
		fmt.Println(ln)      // print ln in main prompt
		// fmt.Println(conn, "I heard you type %s \n", ln) print memory of conn
		fmt.Fprintf(conn, "You typed: %s \n", ln) // shows whats printed
		conn.Close()
	}
	defer conn.Close()

	fmt.Println("ending.....")
}
