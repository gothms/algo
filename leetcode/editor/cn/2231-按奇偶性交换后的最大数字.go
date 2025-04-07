package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 3159
	integer := largestInteger(num)
	fmt.Println(integer)
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestInteger(num int) int {
	// lc：选择排序
	arr := []byte(strconv.Itoa(num))
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if arr[idx] < arr[j] && (arr[j]-arr[idx])&1 == 0 {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
	v, _ := strconv.Atoi(string(arr))
	return v

	// 个人
	//temp := [10]int{}
	//eo := make([]bool, 0) // 偶数为 true
	//for ; num != 0; num /= 10 {
	//	v := num % 10
	//	temp[v]++
	//	eo = append(eo, v&1 == 0)
	//}
	//ans := 0
	//for i, e, o := len(eo)-1, 8, 9; i >= 0; i-- {
	//	if eo[i] { // 偶数
	//		for temp[e] == 0 {
	//			e -= 2
	//		}
	//		ans = ans*10 + e
	//		temp[e]--
	//	} else {
	//		for temp[o] == 0 {
	//			o -= 2
	//		}
	//		ans = ans*10 + o
	//		temp[o]--
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
