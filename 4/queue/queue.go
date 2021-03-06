package queue

// An FIFO queue
type Queue []int //定义别名

// Pushes the element into the queue
func (q *Queue) Push(v int){
	*q = append(*q, v)
}

func (q *Queue) Pop() int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}