package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func generateRing(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	tail := head
	for i := 1; i < len(arr); i++ {
		tail.Next = &ListNode{Val: arr[i]}
		tail = tail.Next
	}
	tail.Next = head
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
