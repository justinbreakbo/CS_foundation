package double_pointer

import (
	"fmt"
	"sort"
	"testing"
)

// 题目描述：归并两个有序数组，把归并结果存到第一个数组上。
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	//go sdk 自带切片排序函数（int类型）
	sort.Ints(nums1)
}

//变长应该用切片
func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
