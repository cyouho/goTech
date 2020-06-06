package main

import "fmt"

func main() {
	var num1 int = 10
	fmt.Printf("num1 的地址是 %v\n", &num1)

	var ptr *int = &num1
	fmt.Printf("ptr 的值是 %v\n", ptr)

	// 下面的编译会报错，因为 num1 是一个int类型，
	// 所以 &num1 也不能赋值给 string 类型的指针 ptr2
	// var ptr2 *string = &num1
	// fmt.Printf("ptr2 的值是 %v\n", *ptr2)

	// 通过指针修改内容
	var num2 int = 20
	var ptr2 *int = &num2
	fmt.Printf("num2 的地址是 %v\n", ptr2)
	*ptr2 = 30
	fmt.Printf("num2 的值是 %v\n", num2)
}
