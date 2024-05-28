package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	limit := 4
	q := [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}
	results := queryResults(limit, q)
	fmt.Println(results)

	queries := [][]int{{2, 3, 1}}
	bools := getResults(queries)
	fmt.Println(bools)

	i := test()
	fmt.Println(i)

	prime := 2099999999
	for v := int(math.Sqrt(float64(prime))); v >= 2; v-- {
		if prime%v == 0 {
			fmt.Println("2099999999 not prime!")
		}
	}
	s := "abcab"
	fmt.Println(len(strings.Split(s, "havbc")))
}
func duplicateNumbersXOR(nums []int) int {
	memo := make(map[int]struct{})
	ans := 0
	for _, v := range nums {
		if _, ok := memo[v]; ok {
			ans ^= v
		} else {
			memo[v] = struct{}{}
		}
	}
	return ans
}
func occurrencesOfElement(nums []int, queries []int, x int) []int {
	memo := make([]int, 0)
	for i, v := range nums {
		if v == x {
			memo = append(memo, i)
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q > len(memo) {
			ans[i] = -1
		} else {
			ans[i] = memo[q-1]
		}
	}
	return ans
}
func queryResults(limit int, queries [][]int) []int {
	ans, vis, memo := make([]int, len(queries)), make(map[int]int), make(map[int]int)
	cnt := 0
	for i, q := range queries {
		idx, color := q[0], q[1]
		if c := vis[idx]; c > 0 {
			memo[c]--
			if memo[c] == 0 {
				cnt--
			}
		}
		if memo[color] == 0 {
			cnt++
		}
		memo[color]++
		vis[idx] = color
		ans[i] = cnt
	}
	return ans
}

// 739 / 742 个通过的测试用例

type stNode struct {
	left, right *stNode
	l, r        int
	length      int
}

func update(i int, cur *stNode) {
	if i == cur.l || i == cur.r {
		cur.length = cur.r - cur.l
		return
	}
	if cur.left == nil {
		cur.left = &stNode{l: cur.l, r: i, length: i - cur.l}
		cur.right = &stNode{l: i, r: cur.r, length: cur.r - i}
	} else {
		if i < cur.left.r {
			update(i, cur.left)
		} else if i > cur.right.l {
			update(i, cur.right)
		}
	}
	cur.length = max(cur.left.length, cur.right.length)
}
func query(i int, cur *stNode) int {
	if cur.left == nil {
		return i - cur.l
	}
	if i <= cur.left.r {
		return query(i, cur.left)
	} else {
		ret := cur.left.length
		if i-cur.right.l <= ret {
			return ret
		}
		return max(ret, query(i, cur.right))
	}
}
func getResults(queries [][]int) []bool {
	n := len(queries)
	maxR := min(50_000, n*3)
	ans := make([]bool, 0)
	st := &stNode{l: 0, r: maxR, length: maxR}
	for _, q := range queries {
		if len(q) == 2 {
			update(q[1], st)
		} else {
			ans = append(ans, query(q[1], st) >= q[2])
		}
	}
	return ans
}

func test() int {
	//const primeRK = 16777619
	//n := 2_100_000_001
	//n := 100_000_001
	n := 16777619 + 1
	prime := make([]bool, n)
	for i := 2; i < n; i++ {
		if prime[i] {
			continue
		}
		for j := i << 1; j < n; j += i {
			prime[j] = true
		}
	}
	for i := n - 1; ; i-- {
		if !prime[i] {
			return i
		}
	}
}
