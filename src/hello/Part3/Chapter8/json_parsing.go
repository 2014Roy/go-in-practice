package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
}

var JSON = `{"name" : "jack"}`

var ks = []byte(`{"firstname": "jach", "lastname": "chen", "age": 86, "education": [{"school": "zhengzhoudaxue"}, {"school": "shanghaidaxue"}]}`)

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			//递归
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			//递归
			printJSON(u)
		}
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var p Person
	//知道json 数据结构的情况下解析到结构体里面
	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
	//如果不知道数据结构该如何尼？我们可以解析到interface{}
	var f interface{}
	err1 := json.Unmarshal(ks, &f)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	printJSON(f)

	//m := f.(map[string]interface{})
	//fmt.Println(m["firstname"])
}
