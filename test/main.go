package main

import "fmt"

func main() {
	var a int = 5
	var b int = 6
	fmt.Println("a= , b= ", a, b)

	a = a + b
	b = a - b
	a = a - b
    fmt.Println("a= , b= ",a ,b)
}