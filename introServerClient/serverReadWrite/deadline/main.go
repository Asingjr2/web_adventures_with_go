package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	// SetDeadline can be be based on time
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("TIMEOUT!!!")
	}
	scanner := bufio.NewScanner(conn)
	// runs forever so connection would not time out without deadline
	for scanner.Scan() {
		ln := scanner.Text()                      // reads information in connection
		fmt.Println(ln)                           // print ln in main prompt
		fmt.Fprintf(conn, "You typed: %s \n", ln) // shows whats printed
	}
	defer conn.Close()

	fmt.Println("ending.....") // Get to this because connection times out
}
