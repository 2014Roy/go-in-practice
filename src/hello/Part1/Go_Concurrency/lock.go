package main

import (
	"fmt"
	"time"
)

func main() {
	//有1个缓存空的channel
	lock := make(chan bool, 1)
	for i := 1; i < 6; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releaseing the lock\n", id)
	<-lock
}
