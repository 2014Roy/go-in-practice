package main

import "net/http"

func main() {
	//Strip- Prefix is used to remove any prefix in the URL before passing the path to the file server to find
	dir := http.Dir("./files/")
	//Strip- Prefix is used to remove any prefix in the URL before passing the path to the file server to find
	handle := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handle)

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	//do something
}
