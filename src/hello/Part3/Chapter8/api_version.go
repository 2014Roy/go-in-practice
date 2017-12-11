package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type textMessage struct {
	Message string `json:"message"`
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	data := textMessage{"jack is coming"}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "json parse error", 500)
		return
	}
	//api version in custom content type
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	//api version in path
	/*
		http.HandleFunc("/api/v1/test", displayTest)
		http.ListenAndServe(":8080", nil)
	*/

	//api version in content type
	http.HandleFunc("/test", displayTest2)
	http.ListenAndServe(":8080", nil)
}

type testMessageV1 struct {
	Message string `json:"message"`
}

type testMessageV2 struct {
	Info string `json:"info"`
}

//api version in custom content type
func displayTest2(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("Accept")
	var err error
	var b []byte
	var ct string
	switch t {
	case "application/vnd.mytodos.json; version = 2.0":
		data := testMessageV2{"Version 2"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version = 2.0"
	case "applicationa/vnd.mytodos.json; version = 1.0":
		fallthrough
	default:
		data := testMessageV1{"Version 1"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version = 1.0"
	}

	if err != nil {
		http.Error(w, "internal server error", 500)
	}
	w.Header().Set("Content-type", ct)
	fmt.Fprint(w, string(b))
}
