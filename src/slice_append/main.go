package main

import "fmt"

func main() {
	// s1 := []string{"bj", "cq"}
	// // 调用append 函数必须用原来的切片变量接收返回值
	// s1 = append(s1, "gc")
	// fmt.Println(len(s1), cap(s1))
	// fmt.Println(s1)
	a1 := [...]int{1, 3, 4, 5, 6, 7, 8, 9, 12}
	s1 := a1[:]
	// 删除索引位1的
	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}
