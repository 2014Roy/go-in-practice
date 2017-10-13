package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("trapped panic: %s (%T) \n", err, err)
		}
	}()

	yikes()
}

func yikes() {
	panic(errors.New("something bad happened"))
}
