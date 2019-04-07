package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string{  //实现Get方法就是实现接口
	return url
}