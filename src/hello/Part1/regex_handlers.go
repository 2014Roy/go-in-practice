package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	//Gets an instance of a path-based router
	rr := newPathResolver()
	rr.Add("GET /hello", hello)
	rr.Add("* /goodbye/*", goodbye)
	//Sets the HTTP server D to use your router
	http.ListenAndServe(":8080", rr)
}

//Creates new initialized pathResolver
func newPathResolver() *pathResolver {
	return &pathResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

//Adds paths to internal lookup
func (r *pathResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

//Iterates over registered paths
func (r *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//checks whether current path matches a registered one
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			//executes the handler function for a matched path
			handlerFunc(res, req)
			return
		}
		http.NotFound(res, req)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Roy"
	}
	fmt.Fprint(res, "hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Roy"
	}
	fmt.Fprint(res, "goodbye  ", name)
}
