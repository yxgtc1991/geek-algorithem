package stack_queue

import (
	"fmt"
	"testing"
)

// MyQueue 232、用栈实现队列
type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// 辅助函数：将inStack所有元素迁移至outStack，保证FIFO
func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		n := len(q.inStack)
		q.outStack = append(q.outStack, q.inStack[n-1])
		q.inStack = q.inStack[:n-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	// 当outStack不为空时，先消耗outStack元素，不需要迁移
	n := len(q.outStack)
	item := q.outStack[n-1]
	q.outStack = q.outStack[:n-1]
	return item
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func TestMyQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}
