package queue

// An FIFO queue
type Queue []int

// push the element into the queue
//   e.g. q.Push(123)
func (q *Queue) Push(value int)  {
	*q = append(*q, value)
}

// pop the element from the queue
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}


// return if the queue is empty or not
func (q *Queue) Empty()bool  {
	return len(*q) == 0
}
