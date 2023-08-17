package stack_queue

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

// isValid 20、有效的括号
func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		// 左括号
		if _, ok := pairs[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			// 右括号
			n := len(stack)
			if n == 0 || stack[n-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:n-1]
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	c.Convey("测试-有效括号匹配", t, func() {
		c.So(isValid("[[[]]]"), c.ShouldEqual, true)
	})
	c.Convey("测试-无效括号匹配", t, func() {
		c.So(isValid("[[]{}{"), c.ShouldEqual, false)
	})
}
