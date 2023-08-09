package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	head := &ListNode{
		Val: 1,
	}
	node1 := &ListNode{
		Val: 2,
	}
	head.Next = node1
	node2 := &ListNode{
		Val: 3,
	}
	node1.Next = node2
	tail := &ListNode{
		Val: 4,
	}
	node2.Next = tail
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
	newHead := reverseList(head)
	curr = newHead
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
