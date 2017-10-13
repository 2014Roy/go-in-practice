package main

import (
	"log"
	"os"
)

func main() {
	//创建错误日志！！
	logfile, _ := os.Create("log.txt")
	defer logfile.Close()

	logger := log.New(logfile, " example", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
	logger.Println("This is the end of the function.")
}
