package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string{
	g := ""
	switch  {  //switch后可以没有表达式
	case score < 60:
		g = "F"
		fallthrough
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g= "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))  //中断程序执行，报错
	}
	return g
}

func main() {
	const filename = "2/abc.txt"
	if contents, err  := ioutil.ReadFile(filename);err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("%s\n", contents)
	} //此种方法在if外不能使用
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println("%s\n", contents)
	//}
	fmt.Println(grade(55))
}
