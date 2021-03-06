package main

import "fmt"

// vscode 不支持go module
func main() {
	// 1. &: 去地址
	// n := 18
	// p := &n
	// fmt.Println(p)
	// fmt.Printf("%T\n", p)  // *int：int类型的指针
	// // 2. *: 根据地址取值
	// m := *p
	// fmt.Println(m)
	// fmt.Printf("%T\n", m)
	var a1 *int // nil pointer
	fmt.Println(a1)
	var a2 = new(int) // new 函数申请一个内存地址
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}
