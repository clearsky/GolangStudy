package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right * Node
}

func createNode(value int) * Node{
	return &Node{Value: value} //局部变量地址也可以使用
}

func (node Node) Print(){  //给struct定义方法
	fmt.Println(node.Value)  //参数传递方式为传值
}

func (node *Node) SetValue(value int){
	node.Value = value
}


