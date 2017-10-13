package main

import "fmt"
import "log"

func main() {
	fmt.Println("vim-go")
	log.Println("this is a regular msg")
	log.Fatalln("this is a fatal error")
	log.Println("this is the end of function")
}
