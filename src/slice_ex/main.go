package main

import (
	"fmt","sort"
)


func main() {
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(cap(a))
	var a1 = [...]int{1, 3, 5, 2, 45, 1, 3}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)
}
