package main

//Using a close channel
import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)
	//Adds an additional Boolean channel that indicates when you’re finished
	done := make(chan bool)

	go send(msg, done)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
		fmt.Println("here can perform ?")
	}
}

//注意区分两个参数类型，第一个是只写通道，第二个是只读通道。
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello boy"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
