package main

import "fmt"
import (
	"errors"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	args := os.Args[1:]

	if result, err := concat(args...); err != nil {
		fmt.Printf("Error : %s\n ", err)
	} else {
		fmt.Printf(" concatenated string: '%s' \n", result)
	}
}

func concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings")
	}

	return strings.Join(parts, "+"), nil
}
