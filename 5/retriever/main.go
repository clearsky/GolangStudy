package main

import (
	"GolangStudy/5/retriever/mock"
	"GolangStudy/5/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever  // interface的组合
	Poster
}

func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}


func main() {
	var r Retriever = mock.Retriever{"this is a fake imooc.com"} //只要实现了interface的规定,就行
	fmt.Println(download(r))
	userAgent := "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 SE 2.X MetaSr 1.0"
	var r2 Retriever = &real.Retriever{userAgent, 5}
	fmt.Println(download(r2))

	fmt.Printf("%T %v", r2, r2)  //type和value
	fmt.Println()
	hehe := r2.(*real.Retriever)
	fmt.Println(hehe.UserAgent)
}
