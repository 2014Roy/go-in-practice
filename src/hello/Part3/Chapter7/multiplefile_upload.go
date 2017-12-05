package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file-upload.html")
		t.Execute(w, nil)
	} else {
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		data := r.MultipartForm
		files := data.File["file"]
		for _, fh := range files {
			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			filename := "/tmp/" + fh.Filename
			fmt.Print(filename)
			out, err := os.Create(filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				panic(err)
				return
			}

			_, err = io.Copy(out, f)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
		}

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
