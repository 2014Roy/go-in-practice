package main

//Hello World web server:
import (
	"fmt"
	"net/http"
)

//Handler for an HTTP request
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo MMM")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}
