package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//get
	fmt.Println("vim-go")
	//设置超时1秒
	cc := &http.Client{Timeout: time.Second}
	res, err := cc.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)

	//delete
	req, _ := http.NewRequest("DELETE", "http://example.com/foo/bar", nil)
	res1, _ := http.DefaultClient.Do(req)
	fmt.Printf("%s", res1.Status)

}
