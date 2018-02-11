package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 1)
	go func(a int) {
		ch <- a
	}(1)
	time.Sleep(2e9)
	select {
	case a := <-ch:
		fmt.Println(a)
	default:
		fmt.Println("err")
	}
}
