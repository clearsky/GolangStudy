package main

import "fmt"

func updateSlice(s []int){
	s[0] = 100  //slice可以改变原值
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 9, 7}
	s := arr[2:6]
	fmt.Println(s)
	updateSlice(s)
	fmt.Println(arr)
	fmt.Println("Reslice:")
	s2 := s[1:3]
	fmt.Println(s2)
	fmt.Println("Extending slice")
	s1 := arr[2:6]
	s3 := s1[3:6]  //应取不到第五个，去最原始的arr中查找
	fmt.Println(s1, s3)
	fmt.Println(cap(s3))
	s4 := append(s3, 1010)  // 添加元素
	fmt.Println(arr)
	s5 := append(s4, 11)  //长度超越cap，系统创建新的数组
	fmt.Println(arr)
	s6 := append(s5,12)
	fmt.Println(s6)
}
