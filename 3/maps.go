package main

import "fmt"

func main() {
	m := map[string]string{  //无序的
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) //m2=empty map
	var m3 map[string]int //m3 = nil
	fmt.Println(m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m{
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	erroTest, ok:= m["dsafdasf"] //取不到不会报错，返回空串
	fmt.Println(erroTest, ok)
	if causeNmae, ok := m["cause"]; ok{
		fmt.Println(causeNmae)
	}else{
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting values")
	delete(m, "name")//key不存在不会报错
	fmt.Println(m)


}
