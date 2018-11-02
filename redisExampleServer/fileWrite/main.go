// Create file using os methods
package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "Action Log"
	actionLog, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	actionLog.Write([]byte("added first \n"))
	actionLog.Write([]byte("added second"))
}
