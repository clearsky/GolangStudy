package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3  //函数外不能使用:=定义变量,
var ss ="kkk" //包内部变量，不是全局变量
var bb = true

var (  //可以避免重复的写var
	cc = 3
	dd = 6
)

func variableZeroValue(){
	var a int
	var s string // 变量名在类型之前
	fmt.Printf("%d %q\n", a, s)  //自动赋初值
}

func variableIntiaVlue(){
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func varibleTypeDecuction(){

	var a, b, c, s = 3, 4, true, "abc"  //不写类型可以同时生成多个不同类型的变量
	fmt.Println(a, b, c, s)
}

func varibleShorter(){
	a, b, c, s := 3, 4, true, "abc"  // :=代表定义变量
	fmt.Println(a, b, c, s)
}

func euler(){
	fmt.Println(
		//cmplx.Pow(math.E, 1i * math.Pi) + 1
		cmplx.Exp(1i * math.Pi) + 1)
}

func triangle(){
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))  //必须手动转
	fmt.Println(c)
}

func consts(){
	const filename = "abc.txt"
	const a, b = 3, 4
	// 常量数值可以做各种类型是哟个
	//const (  也可以使用这种方法
	//
	//)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums(){  //go语言的枚举类型
	const(
		cpp = iota  //自增值
		java
		python
		golang
	)
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello word")
	variableZeroValue()
	variableIntiaVlue()
	varibleTypeDecuction()
	varibleShorter()
	fmt.Println(aa, bb, cc, dd, ss)
	euler()
	triangle()
	consts()
	enums()
}
