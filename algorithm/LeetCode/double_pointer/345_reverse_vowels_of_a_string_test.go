package double_pointer

import (
	"fmt"
	"testing"
)

//题目描述:翻转字符串中的元音字符

// (go的数据类型里面有大学问)
const A = byte('A')
const a = byte('a')
const E = byte('E')
const e = byte('e')
const I = byte('I')
const i = byte('i')
const O = byte('O')
const o = byte('o')
const U = byte('U')
const u = byte('u')

func vowel(s byte) bool {
	switch s {
	case A, a, E, e, I, i, O, o, U, u:
		return true
	}
	return false
}

func reverse(s string) string {
	l := 0
	v := []byte(s)
	r := len(s) - 1 //[l...r]
	for {
		for l < r {
			if vowel(v[l]) {
				break
			}
			l++
		}
		for l < r {
			if vowel(v[r]) {
				break
			}
			r--
		}
		if l >= r {
			break
		}
		v[l], v[r] = v[r], v[l]
		l++
		r--
	}
	return string(v)
}

func TestReverse(t *testing.T) {
	s1 := "leetcode"
	fmt.Println("leetcode reverse:", reverse(s1))
	s2 := "Hello"
	fmt.Println("Hello reverse:", reverse(s2))
}
