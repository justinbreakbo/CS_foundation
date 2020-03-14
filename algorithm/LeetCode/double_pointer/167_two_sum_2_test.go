package double_pointer

import (
	"fmt"
	"testing"
)

//题目描述：在有序数组中找出两个数，使它们的和为 target。

// go里面数组不可变长,用切片来做
func twoSum(numbers []int, target int) (index1, index2 int) {
	index1 = 0
	index2 = len(numbers) - 1
	for index1 < index2 {
		sum := numbers[index1] + numbers[index2]
		if target == sum {
			return index1, index2
		} else if target > sum {
			index1 += 1
		} else {
			index2 -= 1
		}
	}
	return -1, -1
}

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15} //切片
	target := 9
	if index1, index2 := twoSum(numbers, target); index2 < 0 {
		fmt.Print("不存在两个数之和等于", target)
	} else {
		fmt.Print(index1, index2)
	}

}
