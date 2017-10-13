package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("simple.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "an example",
		Content: "hava fun dafsfsafasf",
	}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		fmt.Fprint(w, "a error occoured")
		return
	}
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
