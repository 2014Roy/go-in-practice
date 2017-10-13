package main

import "fmt"

func main() {
	defer sayGoodBye()
	fmt.Println("say hello")
}

func sayGoodBye() {
	fmt.Println("good bye")
}
