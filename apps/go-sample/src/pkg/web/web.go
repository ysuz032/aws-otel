package web

import (
	"fmt"
	"net/http"
)

// HelloWorldHandler returns "Hello, World!" as a response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
