package priority_queue

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// KthLargest 703、数据流中的第 K 大元素
type KthLargest struct {
	// KthLargest 继承自 sort.IntSlice
	// sort.IntSlice 实现 sort.Interface 接口 ，必须实现 Len，Less，Swap 方法
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

// KthLargest 实现 Push 和 Pop 方法，从而实现 heap.Interface 接口，实现小顶堆
func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	arr := kl.IntSlice
	val := arr[len(arr)-1]
	kl.IntSlice = arr[:len(arr)-1]
	return val
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestKthLargest(t *testing.T) {
	nums := []int{5, 3, 2}
	kl := Constructor(2, nums)
	fmt.Println(kl.IntSlice)
	kth := kl.Add(8)
	fmt.Println(kl.IntSlice)
	fmt.Println(kth)
}

func TestSortSlice(t *testing.T) {
	nums := []int{1, 3, 2}
	sorted := sort.IntSlice{}
	sorted = append(sorted, nums...)
	fmt.Println(sorted)
	sorted.Sort()
	fmt.Println(sorted)
}
