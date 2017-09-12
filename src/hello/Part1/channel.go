package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c //wait for value to come in
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) //a channel is created
	a := []int{8, 5, 1, 0, -1}
	go printCount(c) // starts the goroutine
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
