package main

import (
	"fmt"
	"testing"
)

// reverseList 206、反转链表
func reverseList(head *ListNode) *ListNode {
	var prev, curr, next *ListNode = nil, head, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func TestReverseList(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	head := generateLinkedList(arr)
	fmt.Print("反转前：")
	show(head)
	newHead := reverseList(head)
	fmt.Print("反转后：")
	show(newHead)
}
