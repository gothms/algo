package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{10, 2}                       // 210
	nums = []int{3, 30, 34, 5, 9, 92}          // 992534330
	nums = []int{432, 43243}                   // 43243432
	nums = []int{111311, 1113}                 // 1113111311
	nums = []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1} // 9876543210
	nums = []int{0, 0}                         // 0
	number := largestNumber(nums)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {

}

//leetcode submit region end(Prohibit modification and deletion)

func largestNumber_(nums []int) string {
	// 个人
	arr := make([][2]int, len(nums))
	for i, v := range nums { // 每个元素的 pow10
		arr[i][0] = v
		pow := 10
		for v /= 10; v != 0; v /= 10 { // 先 v /= 10，避免 v=0
			pow *= 10
		}
		arr[i][1] = pow
	}
	sort.Slice(arr, func(i, j int) bool { // 排序的算法
		return arr[i][0]*arr[j][1]+arr[j][0] > arr[j][0]*arr[i][1]+arr[i][0]
	})
	if arr[0][0] == 0 { // all zeros
		return "0"
	}
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(strconv.Itoa(v[0]))
	}
	return sb.String()

	// lc
	//sort.Slice(nums, func(i, j int) bool {
	//	x, y := nums[i], nums[j]
	//	sx, sy := 10, 10
	//	for sx <= x {
	//		sx *= 10
	//	}
	//	for sy <= y {
	//		sy *= 10
	//	}
	//	return sy*x+y > sx*y+x
	//})
	//if nums[0] == 0 {
	//	return "0"
	//}
	//ans := []byte{}
	//for _, x := range nums {
	//	ans = append(ans, strconv.Itoa(x)...)
	//}
	//return string(ans)
}
