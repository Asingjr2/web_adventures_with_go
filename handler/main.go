// Need to reinforce handler interface concept
// Handler equals type with ServeHttp method that takes response writter and request
package main

import (
	"fmt"
	"net/http"
)

type Samus string

// Below code will serve http response in the way of a unformatted line.
func (s Samus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Example to show this is an interface")
}

func main() {
	var s Samus
	http.ListenAndServe(":8080", s)
}
