package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := openCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer file.Close()

	//do something

}

func openCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()

	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file\n")
		return file, err
	}

	removeEmptyLines(file)

	return file, err
}

func removeEmptyLines(file *os.File) {
	panic(errors.New("failed parse\n"))
}
