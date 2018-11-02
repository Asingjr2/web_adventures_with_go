// Creating simple example of redis key value store
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Creating file to log user input requests
	myLog, err := os.Create("Key-Value Log")
	if err != nil {
		fmt.Println(err)
	}

	// Creating server or returning error.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	// Ensuring close of connection at end of program.
	defer li.Close()
	defer myLog.Write([]byte("Done for today"))

	// Creating infinit loop to accept connections
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Creating go routine to handle every connection requests
		myLog, err := os.Create("Key-Value Log")
		go handle(conn, myLog.Name())
	}
}

func handle(conn net.Conn, file string) {
	defer conn.Close()

	// User Instructions
	io.WriteString(conn, "\r\n KEY VALUE DB\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	// Key value store that will be sent by connection to server
	data := make(map[string]string)
	// Creating scanner object that will reading information sent into/onto connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()     // scanned "token"
		fs := strings.Fields(ln) // results of input separated by spaces

		if len(fs) < 1 { // if nothing was entered
			continue
		}

		// This will be either get, set, delete
		switch fs[0] {
		case "GET":
			k := fs[1]                                   // key is first fields separated value
			v := data[k]                                 // value is value item for k
			fmt.Fprintf(conn, "You requested %s\r\n", v) // Sent into to connection
		case "SET":
			// Must be at least three words for command (e.g. SET), key, value
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Expected value \r\n")
				continue
			}
			k := fs[1]  // key is second word in field separated input
			v := fs[2]  // value is third input in field separated input
			data[k] = v // Adding element to map structure.  Could also sent to a DB or just log
			addition := "Added value of " + data[k] + "to key" + fs[1] + "\n"
			ioutil.WriteFile(file, []byte(addition), 0666)
		case "DEL":
			if len(fs) != 2 {
				fmt.Fprintln(conn, "Expected item to delete \r\n")
				continue
			}
			k := fs[1] // Finding k from split input deleted.
			delete(data, k)
			deletion := "Deleted " + k + "\n"
			ioutil.WriteFile(file, []byte(deletion), 0666)
		default: // used if no above input cases are met.
			fmt.Fprintln(conn, "Invalid command ", fs[0]+"\r\n")
			continue
		}
	}
}
