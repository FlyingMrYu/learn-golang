package main

import "fmt"

// 数组

// 存放元素的容器
// 必须置顶存放的元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]
	fmt.Printf("a1:%T a2:%T\n", a1, a2)
	a1 = [3]bool{true, false, false}
	fmt.Println(a1)
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 2, 5, 45, 6, 8, 5, 5, 5, 5}
	fmt.Println(a10)
	// 初始化方式：根据索引来初始化
	a3 := [5]int{1, 2}
	a4 := [5]int{0: 1, 3: 2}
	fmt.Println(a3)
	fmt.Println(a4)
	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 根据 for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}
	// 多维数组
	// [[1 2],[3 4],[5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	fmt.Println(a11)

	a12 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a12 {
		sum += v
	}
	fmt.Println(sum)
	for i := 0; i < len(a12); i++ {
		for j := i + 1; j < len(a12); j++ {
			if a12[i]+a12[j] == 8 {
				fmt.Printf("(%d,%d,%t)\n", i, j, i)
			}
		}
	}
}
