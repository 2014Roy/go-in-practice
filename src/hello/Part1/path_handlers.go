package main

import (
	"fmt"
	"net/http"
	"path" //Imports the path package to handle URL matches
	"strings"
)

func main() {
	//Gets an instance of a path-based router
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	//Sets the HTTP server D to use your router
	http.ListenAndServe(":8080", pr)
}

//Creates new initialized pathResolver
func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

//Adds paths to internal lookup
func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

//Iterates over registered paths
func (p *pathResolver) ServeHTTP(res htt.ResponseWriter, req *http.Request) {
	//checks whether current path matches a registered one
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			//executes the handler function for a matched path
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
		http.NotFound(res, req)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name == "Roy"
	}
	fmt.Fprint(res, "hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name == "Roy"
	}
	fmt.Fprint(res, "goodbye  ", name)
}
