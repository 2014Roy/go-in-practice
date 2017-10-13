package main

import "fmt"
import "net/http"

//创建go默认文件系统 有go的默认路径

func main() {
	fmt.Println("vim-go")
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
