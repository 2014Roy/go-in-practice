package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//Creates a channel that will receive a message when 30 seconds have elapsed
	done := time.After(30 * time.Second)
	//Makes a new channel for passing bytes from Stdin to Stdout. Because you haven’t specified a size, this channel can hold only one message at a time.
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
			//Because you don’t particularly care about the contents of the message, you don’t assign the received value to a variable. You just read it off the channel, and the select discards the value
		case <-done:
			fmt.Println("time out")
			os.Exit(0)
		}
	}
}

//Takes a write-only channel (chan<-) and sends any received input to that channel
func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
