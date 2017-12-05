package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Incrementally-file-upload.html")
		t.Execute(w, nil)
	} else {
		mr, err := r.MultipartReader()
		if err != nil {
			panic(err)
		}
		values := make(map[string][]string)
		maxValueBytes := int64(10 << 20)
		for {
			//Attempts to read the next part, breaking the loop if the end of the request is reached
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			//Retrieves the name of the form field, continuing the loop if there’s no name
			name := part.FormName()
			if name == "" {
				continue
			}
			////Retrieves the name of the file if one exists
			filename := part.FileName()
			var b bytes.Buffer
			//If there’s no filename, treats it as a text field
			if filename == "" {
				n, err := io.CopyN(&b, part, maxValueBytes)
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}
				maxValueBytes -= n
				if maxValueBytes == 0 {
					msg := "multipart message too large"
					fmt.Fprint(w, msg)
					return
				}
				//Puts the content for the form field into a map for later access
				values[name] = append(values[name], b.String())
				continue
			}

			dst, err := os.Create("/tmp/" + filename)
			defer dst.Close()
			if err != nil {
				return
			}

			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Fprint(w, "upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
