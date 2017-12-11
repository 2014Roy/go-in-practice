package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Error struct {
	//"-" 忽略该key值
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func JSONError(w http.ResponseWriter, e Error) {
	//创建匿名结构体数据
	data := struct {
		Err Error `json:"error"`
	}{e}
	time.Sleep(time.Second)
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     123,
		Message:  "an error occurred",
	}

	JSONError(w, e)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
