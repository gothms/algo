package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 1, 5, 20, 3} // 3
	nums = []int{3, 5}            // 2
	nums = []int{2, 10, 8}        // 3
	nums = []int{10, 4, 3}        // 2
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

//func minimumDeviation(nums []int) int {
//	// 1.证明：maxV 是结果序列中的一个数
//	// 如果结果序列全为偶数，那么将所有数除以 2 后，答案不会更劣。因此结果中一定会有一个奇数
//	// 假设所有元素对应的奇数中，最大奇数是 maxV，那么结果中的那个奇数一定 <= maxV
//	// 因此将 maxV 对应的数 2^x * maxV 变成 maxV，答案不会更劣
//	// 2.那么剩下的数应“尽可能”和 maxV 接近
//	// 每个 v <= maxV 应该保持 v 不变
//	// 每个 v > maxV 应该保持 v >= maxV 的情况下，尽量对 v/2，但最优解可能需要对某些 v > maxV 的 v 再除一次 2
//	// 3.现在的序列排序后为 v0,v1,...,maxV,vk,...,vn-1
//	// 对于 i >= k 都有 vi/2 < maxV < vi
//	var maxV, minV int
//	for _, v := range nums {
//		v >>= bits.TrailingZeros(uint(v))
//		maxV = max(maxV, v) // 最大的奇数
//	}
//	//minV = maxV
//	//bigV := make([]int, 0)
//	//for _, v := range nums {
//	//	if v&1 == 1 { // 先统一为偶数，再操作。且奇数只有两个可选值
//	//		v <<= 1
//	//	}
//	//	for v >= maxV<<1 { // 尽可能接近 maxV
//	//		v >>= 1
//	//	}
//	//	if v >= maxV {
//	//		bigV = append(bigV, v) // bigV 中都是 >maxV 的偶数，或 =maxV 的奇数
//	//	}
//	//	minV = min(minV, v)
//	//}
//	minV = maxV // 优化：位运算
//	bigV := []int{maxV}
//	k := bits.Len(uint(maxV))
//	for _, v := range nums {
//		if v&1 == 1 {
//			v <<= 1
//		}
//		if v > maxV {
//			v >>= bits.Len(uint(v)) - k // v 现在和 maxV 的二进制位数相等
//			if v < maxV {
//				v <<= 1
//			}
//			bigV = append(bigV, v)
//		}
//		minV = min(minV, v) // 最小的偶数 / maxV
//	}
//	//sort.Ints(bigV) // 排序后更好计算
//	//ans := bigV[len(bigV)-1] - minV
//	//for i := len(bigV) - 1; bigV[i] > maxV; i-- { // maxV 本身也在 bigV 中，所以不会越界
//	//	minV = min(minV, bigV[i]>>1)   // 更新最小值，且 bigV[i] 必然是偶数
//	//	ans = min(ans, bigV[i-1]-minV) // 必有 i > 0
//	//}
//	sort.Sort(sort.Reverse(sort.IntSlice(bigV))) // 优化：倒序排序
//	ans := bigV[0] - minV
//	for i := 0; bigV[i] > maxV; i++ {
//		minV = min(minV, bigV[i]>>1)
//		ans = min(ans, bigV[i+1]-minV)
//	}
//	return ans
//
//	// 堆
//	//minV := math.MaxInt32
//	//h := &hp1675{}
//	//for _, v := range nums {
//	//	if v&1 == 1 { // 全部转为偶数再入堆
//	//		v <<= 1
//	//	}
//	//	minV = min(minV, v)
//	//	heap.Push(h, v) // 大顶堆
//	//}
//	//ans := h.IntSlice[0] - minV
//	//for top := h.IntSlice[0]; top&1 == 0; {
//	//	minV = min(minV, top>>1) // 更新最小值
//	//	h.IntSlice[0] = top >> 1
//	//	heap.Fix(h, 0) // 偶数 /2 后再入堆
//	//
//	//	top = h.IntSlice[0]
//	//	ans = min(ans, top-minV) // 更新最小偏移量
//	//}
//	//return ans
//}
//
//
//type hp1675 struct {
//	sort.IntSlice
//}
//
//func (h hp1675) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
//func (h *hp1675) Push(x any)        { h.IntSlice = append(h.IntSlice, x.(int)) }
//func (h *hp1675) Pop() any {
//	v := h.IntSlice[len(h.IntSlice)-1]
//	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
//	return v
//}
//
//var _ heap.Interface = (*hp1675)(nil)
