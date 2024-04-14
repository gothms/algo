package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	s := "hello"
	ofString := scoreOfString(s)
	fmt.Println(ofString)

	points := [][]int{
		{2, 1},
		{1, 0},
		{1, 4},
		{1, 8},
		{3, 5},
		{4, 6}}
	w := 1
	points = [][]int{{2, 3}, {1, 2}}
	w = 0
	points = [][]int{{1, 3}, {7, 3}}
	w = 1
	points = [][]int{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6}}
	w = 2
	coverPoints := minRectanglesToCoverPoints(points, w)
	fmt.Println(coverPoints)

	arr := [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}
	disappear := []int{1, 1, 5}
	n := 3
	//arr = [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}
	//disappear = []int{1, 3, 5}
	//n = 3
	//arr = [][]int{{0, 1, 1}}
	//disappear = []int{1, 1}
	//n = 2
	time := minimumTime(n, arr, disappear)
	fmt.Println(time)

	nums := []int{1, 4, 3, 3, 2}
	subarrays := numberOfSubarrays(nums)
	fmt.Println(subarrays)
}
func scoreOfString(s string) int {
	sum := 0
	for i := len(s) - 2; i >= 0; i-- {
		if d := int(s[i+1]) - int(s[i]); d > 0 {
			sum += d
		} else {
			sum -= d
		}
	}
	return sum
}
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cnt, n := 0, len(points)
	for i := 0; i < n; {
		cnt++
		last := points[i][0]
		for i < n && last+w >= points[i][0] {
			i++
		}
	}
	return cnt
}
func minimumTime(n int, edges [][]int, disappear []int) []int {
	adj := make([][][2]int, n)
	distant := make([]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]int{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], [2]int{e[0], e[2]})
	}
	h := &mtHp{{}}
	ret := make([]int, n)
	for i := 1; i < n; i++ {
		ret[i], distant[i] = -1, math.MaxInt32
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		if ret[cur[0]] > 0 {
			continue
		}
		ret[cur[0]] = cur[1]
		for _, v := range adj[cur[0]] {
			if dis := cur[1] + v[1]; dis < distant[v[0]] && dis < disappear[v[0]] {
				heap.Push(h, [2]int{v[0], dis})
				distant[v[0]] = dis
			}
		}
	}
	return ret
}

type mtHp [][2]int

func (m mtHp) Len() int           { return len(m) }
func (m mtHp) Less(i, j int) bool { return m[i][1] < m[j][1] }
func (m mtHp) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *mtHp) Push(x any)        { *m = append(*m, x.([2]int)) }
func (m *mtHp) Pop() any {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}
func numberOfSubarrays(nums []int) int64 {
	st := make([][2]int, 0)
	var ret int64
	for _, v := range nums {
		for len(st) > 0 && v > st[len(st)-1][0] {
			k := int64(st[len(st)-1][1])
			ret += k*(k-1)>>1 + k
			st = st[:len(st)-1]
		}
		if len(st) > 0 && v == st[len(st)-1][0] {
			st[len(st)-1][1]++
		} else {
			st = append(st, [2]int{v, 1})
		}
	}
	for _, v := range st {
		k := int64(v[1])
		ret += k*(k-1)>>1 + k
	}
	return ret
}
