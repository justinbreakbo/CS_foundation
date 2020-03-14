package double_pointer

import (
	"fmt"
	"math"
	"testing"
)

//题目描述：判断一个非负整数是否为两个整数的平方和。

func sum(target int) bool {
	if target < 0 {
		return false
	}
	front := 0
	back := int(math.Sqrt(float64(target)))
	for front < back {
		sum := front*front + back*back
		if target == sum {
			return true
		} else if target > sum {
			front += 1
		} else {
			back -= 1
		}
	}
	return false
}

func TestSum(t *testing.T) {
	target := 6
	if ok := sum(target); ok {
		fmt.Println(target, "是两个整数的平方和")
	} else {
		fmt.Println(target, "不是两个整数的平方和")
	}
}
