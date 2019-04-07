package main

import (
	"fmt"
)

func printSlice(s []int){
	fmt.Printf("len=%d, cpa=%d\n", len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	var s []int //Zero value for slice is nil
	for i := 0; i <= 100; i ++{
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	fmt.Println(s1)

	s2 := make([]int, 16, 32)  // 自动赋初值0
	fmt.Println(s2)

	fmt.Println("Copying slice")
	copy(s2, s1)  //覆盖拷贝
	fmt.Println(s2)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) //...相当于python的*
	fmt.Println(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2 [: len(s2) - 1]
	fmt.Println(tail)

	fmt.Println(s2)
	printSlice(s2)
}
