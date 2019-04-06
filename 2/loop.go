package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(n int) string{
	result := ""
	for ; n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result  //数字转字符串
	}
	return result
}

func printFile(filename string){
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func forever(){
	for {  //死循环
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBinary(5),
		convertToBinary(13),
		)
	printFile("2/abc.txt")
	//forever()
}
