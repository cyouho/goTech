package main

import "fmt"

func Print(a interface{}) {
	fmt.Println(a)
}

func main() {
	Print(12)
	Print("hello world")
	Print(12.34)
}