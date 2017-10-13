package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1092")
	if err != nil {
		panic("failed to connect to localhost:1092")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	logger.Println("this is a regular message")
	//使用log的panic可以调用defer函数，如果使用fatal则系统调用os.exit直接终止程序！！！
	logger.Panic("this is a panic")

}
