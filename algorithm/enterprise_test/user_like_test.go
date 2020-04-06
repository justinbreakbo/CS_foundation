package enterprise_test

import "testing"

// 字节跳动原题
// 题目描述：为了不断优化推荐效果，今日头条每天要存储和处理海量数据。
// 假设有这样一种场景：我们对用户按照它们的注册时间先后来标号，对于一类文章，每个用户都有不同的喜好值，
// 我们会想知道某一段时间内注册的用户（标号相连的一批用户）中，有多少用户对这类文章喜好值为k。
// 因为一些特殊的原因，不会出现一个查询的用户区间完全覆盖另一个查询的用户区间(不存在L1<=L2<=R2<=R1)。

func userLike(n int, likeLevel []int, q int, query [][]int) []int {
	var counts, temp []int
	var count int
	for i := 0; i < q; i++ {
		for j := 0; j < n; j++ {
			if likeLevel[j] >= query[i][0] && likeLevel[j] <= query[i][1] {
				temp = append(temp, likeLevel[j])
			}
		}
		count = 0
		for k := 0; k < len(temp); k++ {
			if temp[k] == query[i][2] {
				count++
			}
		}
		counts = append(counts, count)
		temp = nil
	}
	return counts
}

func TestUserLike(t *testing.T) {
	n := 5
	likeLevel := []int{1, 2, 3, 3, 5}
	q := 3
	query := [][]int{{1, 2, 1}, {2, 4, 5}, {3, 5, 3}}
	var counts []int
	counts = userLike(n, likeLevel, q, query)
	for i := 0; i < len(counts); i++ {
		println(counts[i])
	}
}
