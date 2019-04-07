package main

import (
	"GolangStudy/4"
	"GolangStudy/4/queue"
	"fmt"
)

type myTreeNode struct {  //使用组合
	node *tree.Node
}
func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(root)
	fmt.Println(nodes)
	root.SetValue(6) //自动传入指针
	root.Print()
	root.Traverse()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
