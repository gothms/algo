package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	//arr := []int{72, 48, 24, 3}
	//pairs := countCompleteDayPairs(arr)
	//fmt.Println(pairs)

	//arr := []int{7, 1, 6, 6}
	//arr = []int{7, 1, 6, 3} // 10
	//damage := maximumTotalDamage(arr)
	//fmt.Println(damage)

	nums := []int{3, 1, 4, 2, 5}
	queries := [][]int{{2, 3, 4}, {1, 0, 4}}
	nums = []int{4, 1, 4, 2, 1, 5}
	queries = [][]int{{2, 2, 4}, {1, 0, 2}, {1, 0, 4}}
	nums = []int{4, 9, 4, 10, 7}
	queries = [][]int{{2, 3, 2}, {2, 1, 3}, {1, 2, 3}}
	nums = []int{5, 4, 8, 6}
	queries = [][]int{{1, 2, 2}, {1, 1, 2}, {2, 1, 6}}
	nums = []int{8, 5, 9, 3, 5}
	queries = [][]int{{1, 2, 4}, {1, 0, 1}, {2, 2, 4}}
	peaks := countOfPeaks(nums, queries)
	fmt.Println(peaks)
}
func countCompleteDayPairs(hours []int) int {
	memo := make(map[int]int)
	ans := 0
	for _, h := range hours {
		h %= 24
		//fmt.Println(h)
		ans += memo[(24-h)%24]
		memo[h]++
	}
	return ans
}

//func countCompleteDayPairs(hours []int) int64 {
//	memo := make(map[int]int)
//	ans := 0
//	for _, h := range hours {
//		h %= 24
//		//fmt.Println(h)
//		ans += memo[(24-h)%24]
//		memo[h]++
//	}
//	return int64(ans)
//}

func maximumTotalDamage(power []int) int64 {
	//sort.Ints(power)
	//memo := make(map[int]int)
	//for _, v := range power {
	//	memo[v]++
	//}
	//power = slices.Compact(power) // 去重
	//n := len(power)
	//dp := make([]int, n+1)
	//for i, v := range power {
	//	j := i
	//	for j > 0 && power[j-1]+2 >= v {
	//		j--
	//	}
	//	dp[i+1] = max(dp[i], dp[j]+v*memo[v])
	//}
	//return int64(dp[n])

	sort.Ints(power)
	memo := make(map[int]int)
	for _, v := range power {
		memo[v]++
	}
	power = slices.Compact(power)
	dp := make([]int, len(power))
	dp[0] = power[0] * memo[power[0]]
	for i := 1; i < len(power); i++ {
		v := power[i]
		j := i - 1
		dp[i] = max(dp[j], v*memo[v])
		if power[j]+2 < v {
			dp[i] = dp[j] + v*memo[v]
			continue
		}
		if j--; j < 0 {
			continue
		} else if power[j]+2 < v {
			dp[i] = max(dp[i], dp[j]+v*memo[v])
			continue
		}
		if j--; j < 0 {
			continue
		} else {
			dp[i] = max(dp[i], dp[j]+v*memo[v])
		}
	}
	return int64(dp[len(power)-1])
}
func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)
	bit, memo := make([]int, n+1), make([]int, n)
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			memo[i] = 1
			update(bit, i+1, 1)
		}
	}
	ans := make([]int, 0)
	for _, q := range queries {
		if q[0] == 1 {
			if q[1]+1 >= q[2] {
				ans = append(ans, 0)
			} else {
				ans = append(ans, rangeSum(bit, q[1]+2, q[2])) // 区间需要去首去尾
			}
		} else {
			index := q[1]
			nums[index] = q[2]
			for i := -1; i <= 1; i++ {
				if idx := index + i; idx > 0 && idx < n-1 {
					v := 0
					if nums[idx] > nums[idx-1] && nums[idx] > nums[idx+1] {
						v = 1
					}
					if delta := v - memo[idx]; delta != 0 { // 需要更新
						update(bit, idx+1, delta)
						memo[idx] = v
					}
				}
			}
		}
	}
	return ans
}
func update(b []int, i, delta int) {
	for ; i < len(b); i += i & -i {
		b[i] += delta
	}
}
func prefixSum(b []int, i int) int {
	sum := 0
	for ; i > 0; i &= i - 1 {
		sum += b[i]
	}
	return sum
}
func rangeSum(b []int, f, t int) int {
	return prefixSum(b, t) - prefixSum(b, f-1)
}
