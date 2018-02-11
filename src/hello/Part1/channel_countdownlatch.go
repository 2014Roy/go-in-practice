package main

import "fmt"

func main() {
	nTask := 5
	ch := make(chan int)
	for i := 1; i <= nTask; i++ {
		go doTask(ch, i)
	}

	for i := 1; i <= nTask; i++ {

		fmt.Printf("管道取值%d\n", <-ch)
	}
	fmt.Println("finished all tasks")
}

func doTask(ch chan int, number int) {
	ch <- number
}
