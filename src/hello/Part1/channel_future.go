package main

import (
	"fmt"
	"time"
)

func main() {
	future := heavyCalculation()
	fmt.Println(<-future)
}

func heavyCalculation() chan int {
	future := make(chan int)
	go func() {
		//模拟耗时操作
		time.Sleep(5e9)
		future <- 666
	}()

	return future
}
