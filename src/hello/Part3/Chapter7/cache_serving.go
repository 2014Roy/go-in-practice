package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

//缓存第一次serverfile 返回的内容到内存

type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile
var mutex = new(sync.RWMutex)

func main() {
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	//读写锁即是针对于读写操作的互斥锁。它与普通的互斥锁最大的不同就是，它可以分别针对读操作和写操作进行锁定和解锁操作。读写锁遵循的访问控制规则与互斥锁有所不同。在读写锁管辖的范围内，它允许任意个读操作的同时进行。但是，在同一时刻，它只允许有一个写操作在进行。并且，在某一个写操作被进行的过程中，读操作的进行也是不被允许的。也就是说，读写锁控制下的多个写操作之间都是互斥的，并且写操作与读操作之间也都是互斥的。但是，多个读操作之间却不存在互斥关系。
	mutex.RLock()
	v, found := cache[req.URL.Path]
	mutex.RUnlock()

	if !found {
		mutex.Lock()
		defer mutex.Unlock()
		fileName := "./files" + req.URL.Path
		//本地判断是否有该文件
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("本地没有 %s \n", fileName)
			http.NotFound(res, req)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			http.NotFound(res, req)
			return
		}
		r := bytes.NewReader(b.Bytes())
		//开始读取操作
		info, _ := f.Stat()
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}
	//缓存提供返回
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
