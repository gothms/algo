package main

import "fmt"

func main() {
	n, k := 3, 2
	n, k = 1, 1
	n, k = 2, 2
	n, k = 10, 4 // 440
	n, k = 10, 5 // 1068
	n, k = 10, 7 // 4489
	inversePairs := kInversePairs(n, k)
	fmt.Println(inversePairs)

	/*
		找到具体的状态转移方程
		3	[1 2 2  1 0 0 0 0 0 0 0]	// 4
		4	[1 3 5  6   5  3  1 0 0 0 0]	// 7
		5	[1 4 9  15 20 22  20  15  9   4  1]	// 11
		6	[1 5 14 29 49 71  90  101 101 90  71  49  29 14 5 1] // 16
			[1 6 20 49 98 169 259 360 461 551 622 671 700 714 719 720]	// 22
		7	[1 6 20 49 98 169 259 359 455 531 573 573 531 455 359 259 169 98 49 20 6 1]
	*/
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = []int{1, 2, 3, 4, 5, 6}
	m := len(arr)
	cnt := make([]int, m*(m-1)>>1+1)
	getCnt := func() {
		c := 0
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				if arr[i] > arr[j] {
					c++
				}
			}
		}
		cnt[c]++
	}
	var dfs func(int)
	dfs = func(i int) {
		if i == m {
			getCnt()
			return
		}
		dfs(i + 1)
		for j := i + 1; j < m; j++ {
			arr[i], arr[j] = arr[j], arr[i]
			dfs(i + 1)
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	dfs(0)
	fmt.Println(cnt, len(cnt))
	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}
	//fmt.Println(cnt)
}

// leetcode submit region begin(Prohibit modification and deletion)
func kInversePairs(n int, k int) int {
	// dp：个人
	if k == 0 { // fast path
		return 1
	} else if n == 1 {
		return 0
	}
	m := n*(n-1)>>1 + 1
	if k >= m { // 超出可有的逆序对数
		return 0
	} else if k == m-1 { // 倒序排序
		return 1
	}
	const mod = 1_000_000_007
	last := (m - 1) >> 1 // n 个数有 m 个逆序对，中位值是 last
	if k > last {        // 超过中位值，则映射到中位值之前
		k = last<<1 - k + (m+1)&1
	}
	dp, temp := make([]int, k+1), make([]int, k+1)
	dp[0] = 1
	for i, maxIP := 2, 1; i <= n; i++ { // cur：最多 maxIP 个逆序对
		for j := 1; j <= min(maxIP, k); j++ {
			//if j > maxIP>>1 { // 也可以通过状态转移方程计算
			//	temp[j] = temp[maxIP-j]
			//	continue
			//}
			if j >= i {
				// - dp[j-i]：比如 [1,2,3,4,5] 求 k = 5
				temp[j] = (temp[j-1] + dp[j] - dp[j-i] + mod) % mod
			} else {
				// 在 0 到 j-1 个逆序对的队列中，适当位置插入数字 j，或 i-1 个数的队列有 j 个逆序对，在末尾追加数字 j
				temp[j] = (temp[j-1] + dp[j]) % mod
			}
		}
		maxIP += i
		dp, temp = temp, dp // 滚动数组
	}
	return dp[k]
}

//leetcode submit region end(Prohibit modification and deletion)
