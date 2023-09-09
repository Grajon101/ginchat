package main

import "fmt"

func main2() {
	// 定义一个字符串类型的变量
	var myAddr = "tree road 1025, 100"

	// 对字符串取地址， ptr类型为*string
	ptr := &myAddr

	// 打印ptr的类型
	fmt.Printf("ptr的类型是 : %T\n", ptr)

	// 打印ptr的指针地址
	fmt.Printf("ptr的地址是 : %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 打印取值后的类型
	fmt.Printf("value的类型是 : %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value的值是 : %s\n", myAddr)

}
