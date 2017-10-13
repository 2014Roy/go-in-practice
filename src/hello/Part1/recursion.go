package main

//递归练习

import "fmt"

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		fmt.Println(result)
		return result
	}

	return 1
}

func main() {
	var i int = 13
	fmt.Println("%d 的阶乘 %d/n", i, factorial(uint64(i)))
}
