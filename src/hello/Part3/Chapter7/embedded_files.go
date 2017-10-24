package main

import (
	//go.rice enables you to work with files from the filesystem during development (for example, when using go run), to use files embedded in the built binary, and to build files into binaries
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	box := rice.MustFindBox("./files/")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}

//$ go get github.com/GeertJohan/go.rice/rice

//After this tool is installed, you can build a Go binary with the following two commands:
// $ rice embed-go
//  $ go build
//	  The first command, rice embed-go, converts the real filesystem elements into a vir- tual filesystem inside Go files. This includes the content of the files
