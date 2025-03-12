package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	// lc
	n := len(flowers)
	left := int(newFlowers - int64(n*target))
	for i, v := range flowers {
		flowers[i] = min(target, v)
		left += flowers[i] // 剩下 left 朵花
	}
	if int64(left) == newFlowers { // 已经全部栽满
		return int64(n * full)
	}
	if left >= 0 { // 可全部栽满
		return max(int64((n-1)*full+(target-1)*partial), int64(n*full))
	}

	sort.Ints(flowers) // 排序
	preSum, j := flowers[0], 1
	//preSum, j := 0, 0
	var ans int64
	for i := 1; i <= n; i++ { // 易错：<=
		left += target - flowers[i-1] // 剩下的花
		if left < 0 {                 // 不够栽
			continue
		}
		for ; j < i && flowers[j]*j <= preSum+left; j++ { // 尝试在 [0,j] 栽上 flowers[j] 朵花
			preSum += flowers[j]
		}
		avg := (preSum + left) / j // 在 [0,j] 平均栽上 avg 朵花
		ans = max(ans, int64((n-i)*full+avg*partial))
	}
	return ans

	// 个人：二分
	//sort.Ints(flowers)
	//n := len(flowers)
	//l, r := 0, n
	//var ans int64
	//for l < r {
	//	m := (l + r) >> 1
	//
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
