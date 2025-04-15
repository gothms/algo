package main

import "fmt"

func main() {
	//fmt.Println(-1 & 1)
	nums := []int{1, 1, 1, 1, 1}
	k := 10
	nums = []int{3, 1, 4, 3, 2, 2, 4}
	k = 2
	good := countGood(nums, k)
	fmt.Println(good)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countGood(nums []int, k int) int64 {
	memo := make(map[int]int)
	ans, cur, i := 0, 0, 0
	for _, v := range nums {
		cur += memo[v]
		memo[v]++
		for ; cur >= k; i++ {
			memo[nums[i]]--
			cur -= memo[nums[i]]
		}
		ans += i
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func countGood_(nums []int, k int) int64 {
	// 双指针
	ret, cnt, n := int64(0), 0, len(nums)
	memo := make(map[int]int)
	for i, j := 0, 0; i < n; i++ {
		cnt += memo[nums[i]]
		memo[nums[i]]++
		for ; cnt >= k; j++ {
			memo[nums[j]]--
			cnt -= memo[nums[j]]
			ret += int64(n - i)
		}
	}
	return ret

	// 优化
	//cg, cnt, l, n := int64(0), 0, 0, len(nums)
	//memo := make(map[int]int)
	//for r, v := range nums {
	//	// x 个相同的数，组成 (x-1)*x / 2 对
	//	cnt += memo[v] // 在 m[v]++ 之前，对数的增量为 m[v]
	//	memo[v]++
	//	for cnt >= k {
	//		memo[nums[l]]--
	//		cnt -= memo[nums[l]] // 对数的增量为 m[v]
	//		l++
	//		cg += int64(n - r) // r 及其后面的元素，都能成为一个好子数组
	//	}
	//}
	//return cg

	// 滑动窗体
	//cg, cnt, l, n := int64(0), 0, 0, len(nums)
	//memo := make(map[int]int)
	//for r, v := range nums {
	//	// x 个相同的数，组成 (x-1)*x / 2 对
	//	cnt += memo[v] // 在 m[v]++ 之前，对数的增量为 m[v]
	//	memo[v]++
	//	if cnt < k {
	//		continue
	//	}
	//	cg += int64(n - r) // r 及其后面的元素，都能成为一个好子数组
	//	for {
	//		memo[nums[l]]--
	//		cnt -= memo[nums[l]] // 对数的增量为 m[v]
	//		l++
	//		if cnt < k {
	//			break
	//		}
	//		cg += int64(n - r)
	//	}
	//}
	//return cg
}
