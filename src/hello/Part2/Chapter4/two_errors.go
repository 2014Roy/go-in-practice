package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("request timed out")
var ErrRejected = errors.New("request was rejected")

var random = rand.New(rand.NewSource(35))

func main() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. retrying")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
