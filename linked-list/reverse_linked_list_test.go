package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 反转链表
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

// generateLinkedList 构造链表
func generateLinkedList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	tail := head
	for i := 1; i < len(arr); i++ {
		tail.Next = &ListNode{Val: arr[i]}
		tail = tail.Next
	}
	return head
}

// show 打印链表
func show(head *ListNode) {
	curr := head
	for curr != nil {
		if curr.Next == nil {
			fmt.Print(curr.Val)
			break
		}
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println()
}

func TestReverseList(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	head := generateLinkedList(arr)
	fmt.Println("反转前：")
	show(head)
	newHead := reverseList(head)
	fmt.Println("反转后：")
	show(newHead)
}
