package main

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

// hasCycle 141、环形链表
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
	c.Convey("测试-有环", t, func() {
		c.So(hasCycle(head), c.ShouldEqual, true)
	})
	head2 := generateLinkedList(arr)
	c.Convey("测试-无环", t, func() {
		c.So(hasCycle(head2), c.ShouldEqual, false)
	})
}
