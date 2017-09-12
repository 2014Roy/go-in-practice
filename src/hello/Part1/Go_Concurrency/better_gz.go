package main

//安全压缩大量文件

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	//A WaitGroup doesn't need to be init
	var wg sync.WaitGroup
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		//开辟线程执行匿名函数
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("compressed %d file\n", i+1)
}

func compress(filename string) error {

}
