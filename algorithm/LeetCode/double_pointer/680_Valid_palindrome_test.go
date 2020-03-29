package double_pointer

import (
	"fmt"
	"testing"
)

//题目描述：给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

func palindrome(s string) bool {
	if s == "" {
		return false
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left += 1
			right -= 1
		} else {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
	}
	return true
}

func isPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] == s[right] {
			left += 1
			right -= 1
		} else {
			return false
		}
	}
	return true
}

func TestPalindrome(t *testing.T) {
	s1 := "aba"
	s2 := "abca"
	if ok := palindrome(s1); ok {
		fmt.Println("s1符合要求")
	} else {
		fmt.Println("s1不符合要求")
	}
	if ok := palindrome(s2); ok {
		fmt.Println("s2符合要求")
	} else {
		fmt.Println("s2不符合要求")
	}
}
