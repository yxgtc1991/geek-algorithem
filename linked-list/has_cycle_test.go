package main

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	// 快慢指针
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func TestHasCycle(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	head := generateRing(arr)
	isCycle := hasCycle(head)
	c.Convey("环形链表", t, func() {
		c.So(isCycle, c.ShouldEqual, true)
	})
}
