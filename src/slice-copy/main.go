package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
}
