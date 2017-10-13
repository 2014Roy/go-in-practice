package main

import (
	"net/http"
)

//When you want to write your own handler to serve files, the ServeFile function in the http package is useful 自定义文件处理function

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/appstore.pdf")
}
func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8080", nil)
}
