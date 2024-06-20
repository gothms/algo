package main

import (
	"container/heap"
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4}                                      // 1
	nums = []int{4, 1, 5, 20, 3}                                   // 3
	nums = []int{3, 5}                                             // 1
	nums = []int{2, 10, 8}                                         // 3
	nums = []int{10, 4, 3}                                         // 2
	nums = []int{399, 908, 648, 357, 693, 502, 331, 649, 596, 698} // 315
	deviation := minimumDeviation(nums)
	fmt.Println(deviation)

	//v := 40
	//v >>= bits.TrailingZeros(uint(v))
	//fmt.Println(v)
}

// https: //leetcode.cn/problems/minimize-deviation-in-array/solutions/503280/yi-chong-fu-za-du-geng-di-de-zuo-fa-by-heltion-2/
// leetcode submit region begin(Prohibit modification and deletion)
func minimumDeviation(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func minimumDeviation_(nums []int) int {
	// 终版
	maxOdd := 0
	for _, v := range nums {
		maxOdd = max(maxOdd, v>>bits.TrailingZeros(uint(v))) // 最大的奇数
	}
	minV, i, k := maxOdd, 0, bits.Len(uint(maxOdd)) // maxOdd 的二进制位数
	for _, v := range nums {
		if v&1 == 1 {
			v <<= 1
		}
		if v > maxOdd {
			v >>= bits.Len(uint(v)) - k // v 现在和 maxOdd 的二进制位数相等
			nums[i] = v
			i++
		}
		minV = min(minV, v) // 最小的数 / maxV
	}
	sort.Ints(nums[:i]) // 排序
	ans := nums[i-1] - minV
	for j := i - 1; nums[j] > maxOdd; j-- {
		minV = min(minV, nums[j]>>1)   // 更新最小值
		ans = min(ans, nums[j-1]-minV) // 更新最小偏移量
	}
	return ans

	// 1.证明：maxOdd 是结果序列中的一个数
	// 如果结果序列全为偶数，那么将所有数除以 2 后，答案不会更劣。因此结果中一定会有一个奇数
	// 假设所有元素对应的奇数中，最大奇数是 maxOdd，那么结果中的那个奇数一定 <= maxOdd
	// 因此将 maxOdd 对应的数 2^x * maxOdd 变成 maxOdd，答案不会更劣
	// 2.那么剩下的数应“尽可能”和 maxOdd 接近
	// 每个 v <= maxOdd 应该保持 v 不变
	// 每个 v > maxOdd 应该保持 v >= maxOdd 的情况下，尽量对 v/2，但最优解可能需要对某些 v > maxOdd 的 v 再除一次 2
	// 3.现在的序列排序后为 v0,v1,...,maxOdd,vk,...,vn-1
	// 对于 i >= k 都有 vi/2 < maxOdd < vi
	//var maxOdd, minV int
	//for _, v := range nums {
	//	v >>= bits.TrailingZeros(uint(v))
	//	maxOdd = max(maxOdd, v) // 最大的奇数
	//}
	//minV = maxOdd
	//bigV := make([]int, 0)
	//for _, v := range nums {
	//	if v&1 == 1 { // 先统一为偶数，再操作。且奇数只有两个可选值
	//		v <<= 1
	//	}
	//	for v >= maxOdd<<1 { // 尽可能接近 maxOdd
	//		v >>= 1
	//	}
	//	if v >= maxOdd {
	//		bigV = append(bigV, v) // bigV 中都是 >maxOdd 的偶数，或 =maxOdd 的奇数
	//	}
	//	minV = min(minV, v)
	//}
	//sort.Ints(bigV) // 排序后更好计算
	//ans := bigV[len(bigV)-1] - minV
	//for i := len(bigV) - 1; bigV[i] > maxOdd; i-- { // maxOdd 本身也在 bigV 中，所以不会越界
	//	minV = min(minV, bigV[i]>>1)   // 更新最小值，且 bigV[i] 必然是偶数
	//	ans = min(ans, bigV[i-1]-minV) // 必有 i > 0
	//}
	//return ans

	// 堆
	//minV := math.MaxInt32
	//h := &hp1675{}
	//for _, v := range nums {
	//	if v&1 == 1 { // 全部转为偶数再入堆
	//		v <<= 1
	//	}
	//	minV = min(minV, v)
	//	heap.Push(h, v) // 大顶堆
	//}
	//ans := h.IntSlice[0] - minV
	//for top := h.IntSlice[0]; top&1 == 0; {
	//	minV = min(minV, top>>1) // 更新最小值
	//	h.IntSlice[0] = top >> 1
	//	heap.Fix(h, 0) // 偶数 /2 后再入堆
	//
	//	top = h.IntSlice[0]
	//	ans = min(ans, top-minV) // 更新最小偏移量
	//}
	//return ans
}

type hp1675 struct {
	sort.IntSlice
}

func (h hp1675) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp1675) Push(x any)        { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp1675) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

var _ heap.Interface = (*hp1675)(nil)
