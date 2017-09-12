package main

//如何优雅的关闭服务，同时避免数据丢失以及做清理工作。利用一个库github.com/braintree/manners.

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

//Handler responding to web requests
func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "inigo mon...."
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

func main() {
	//Gets instance of a handler
	handler := newHandler()
	//Sets up monitoring of operating system signals
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)
	// Starts the web server
	manners.ListenAndServe(":8080", handler)
}
