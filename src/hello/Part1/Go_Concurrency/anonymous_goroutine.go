package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(" outside a goroutine")
	go func() {
		fmt.Println("inside a goroutine")
	}()
	fmt.Println("outside again")

	//让调度器屈服等待goruntine的执行完成
	runtime.Gosched()
}
