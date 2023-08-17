package main

import (
	"fmt"
	"testing"
)

// swapPairs 24、两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	// 哑节点：后继指针指向新头节点
	dummy := &ListNode{0, head}
	var prev, curr, next *ListNode = dummy, nil, nil
	for prev.Next != nil && prev.Next.Next != nil {
		curr = prev.Next
		next = prev.Next.Next
		prev.Next = next
		curr.Next = next.Next
		next.Next = curr
		prev = curr
	}
	return dummy.Next
}

func TestSwapPairs(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	head := generateLinkedList(arr)
	fmt.Print("交换前：")
	show(head)
	newHead := swapPairs(head)
	fmt.Print("交换后：")
	show(newHead)
}
