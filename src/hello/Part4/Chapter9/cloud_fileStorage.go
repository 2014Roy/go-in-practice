package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("vim-go")

	content := `loading ipsum` + `aaaaa` + `bbbb` + `cccc`
	//创建字节数组并初始化内容
	body := bytes.NewReader([]byte(content))
	store, err := fileStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Storing content")
	err = store.Save("foo/bar", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("retrieving content")
	c, err := store.Load("foo/bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	o, err := ioutil.ReadAll(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))

}

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	//文件路径
	fmt.Printf("basepath :%s \n", l.Base)
	p := filepath.Join(l.Base, path)
	fmt.Printf("p :%s \n", p)
	//文件目录
	d := filepath.Dir(p)
	fmt.Printf("d :%s \n", d)
	//文件目录
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

func fileStore() (File, error) {
	return &LocalFile{Base: "."}, nil
}
