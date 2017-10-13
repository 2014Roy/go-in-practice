package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("o 不能作为分母")

func main() {
	fmt.Println("div 1 by o")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Println("div 2 by 0")
	divide(2, 0)
}

//检查分母
func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	return divide(a, b), nil
}

//The regular divide function wraps the division operator with no checks
func divide(a, b int) int {
	return a / b
}
