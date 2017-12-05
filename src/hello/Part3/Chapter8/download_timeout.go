package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
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

func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "btyes="+start+"-")
	}
	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "unsuccess http request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func main() {
	file, err := os.Create("file.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	location := "http://pic4.nipic.com/20091217/3885730_124701000519_2.jpg"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}
