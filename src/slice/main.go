package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"1", "123"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 长度和容量
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	// 有数组得到切片
	a1 := [...]int{1, 23, 4, 5, 6, 7, 3, 3, 3, 6}
	// 基于一个数组切割，左保函 右不包含
	s3 := a1[0:5]
	fmt.Println(s3)
	s4 := a1[1:6]
	s5 := a1[:5]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s4, s5, s6, s7, len(s5))
	// 切片在切片
	s8 := s6[1:]
	fmt.Println(s8)
	// 切片是一个应用类型，都指向底层的一个数组
	a1[5] = 2000
	fmt.Println(s8)
}
