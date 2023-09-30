package main

import "fmt"

func main() {
	edges := [][]int{{0, 1},
		{0, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6}}
	values := []int{3, 0, 6, 1, 5, 2, 1}
	//values = []int{3, 2, 7, 1, 4, 2, 5}
	n := 7
	k := 3
	//edges = [][]int{{0, 2},
	//	{1, 2},
	//	{1, 3},
	//	{2, 4}}
	//values = []int{1, 8, 1, 4, 4}
	//n = 5
	//k = 6
	components := maxKDivisibleComponents(n, edges, values, k)
	fmt.Println(components)
}
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	//adj, visited := make([][]int, n), make([]bool, n)
	cnt, adj := 0, make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	var dfs func(int, int) int
	dfs = func(f, t int) int {
		v := values[t]
		for _, i := range adj[t] {
			if i != f {
				v += dfs(t, i)
			}
		}
		if v %= k; v == 0 {
			cnt++
		}
		return v
	}
	dfs(-1, 0)
	return cnt
}
