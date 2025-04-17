package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := 12   // 21
	n = 21    // -1
	n = 92431 // 93124
	//n = 2147483674 // 93124
	element := nextGreaterElement(n)
	fmt.Println(element)
}

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(n int) int {
	// lc
	nums := []byte(strconv.Itoa(n))
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	j := len(nums) - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	ans, _ := strconv.Atoi(string(nums))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

//func nextGreaterElement_(n int) int {
//	buf := make([]int, 0)
//	cur, last := 0, 0
//	for ; n > 0; n /= 10 {
//		cur = n % 10
//		if cur < last {
//			break
//		}
//		buf = append(buf, cur)
//		last = cur
//	}
//	if n == 0 { // n 是倒序排序的数列组成
//		return -1
//	}
//	i := sort.Search(len(buf), func(i int) bool { // buf 本身是升序
//		return buf[i] > cur
//	})
//	cur, buf[i] = buf[i], cur // 变更大
//	n = n/10*10 + cur
//	for _, v := range buf { // 加上升序的数列
//		n = n*10 + v
//	}
//	if n > math.MaxInt32 {
//		return -1
//	}
//	return n
//}
