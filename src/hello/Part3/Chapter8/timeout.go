package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
		errTxt := "use of closed network connection"
		if err != nil && strings.Contains(err.Error(), errTxt) {
			return true
		}
	}

	return false
}

func main() {
	_, err := http.Get("http://www.baidu.com")
	if err != nil && hasTimedOut(err) {
		fmt.Println("A timeout error occured")
		os.Exit(-1)
	}
}
