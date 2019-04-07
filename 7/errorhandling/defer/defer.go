package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer(){
	defer fmt.Println(1)  //在函数退出打出来
	defer fmt.Println(2)  //先进后出
	fmt.Println(3)
}

func writeFile(filename string){
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL, 0666)
	if err != nil{
		if pathError, ok := err.(*os.PathError); !ok{
			panic(err)
		}else{
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		fmt.Println(err.Error())
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i:=0; i < 20;i++{
		fmt.Fprintln(writer, i)
	}

}

func main() {
	tryDefer()
	writeFile("test.txt")
}
