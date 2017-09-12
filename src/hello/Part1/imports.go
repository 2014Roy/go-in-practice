package main

import (
    "fmt"
    "math"
)

func main () {
    fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
    //首字母大写的可以输出 pi 不可输出
    fmt.Printf(math.Pi)
}
