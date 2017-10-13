package main

import "fmt"
import (
	"html/template"
	"net/http"
)

//模板解析设置为全局 防止每次解析模板
var t = template.Must(template.ParseFiles("simple.html"))

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "an example",
		Content: "have fun .......",
	}
	t.Execute(w, p)
}

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
