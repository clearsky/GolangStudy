package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error){
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)  // 下划线表示不想用
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}
}

func div(a, b int) (q, r int){
	//q = a/ b  另一种方式
	//r = a % b
	//return
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling functions %s with args: %d, %d\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int{
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int{  //可变参数列表
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func swap(a, b *int){
	*b, *a = *a, *b
}

func main() {
	fmt.Println(eval(5, 6, "/"))
	fmt.Println(div(13, 3))
	fmt.Println(apply(pow, 6, 7))
	fmt.Println(apply(  //匿名函数方式
		func(a int, b int) int{
			return a + b
		}, 88, 66),
		)
	fmt.Println(sum(5, 6, 7, 8, 9))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
