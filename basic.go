package main

import "fmt"

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

func main() {
	fmt.Println("hello word")
	variableZeroValue()
	variableIntiaVlue()
	varibleTypeDecuction()
	varibleShorter()
	fmt.Println(aa, bb, cc, dd, ss)
}
