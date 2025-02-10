package main

import "fmt"

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	//arr = []int{1, 0, 2, 3, 0, 0, 5, 0}
	//arr = []int{1, 5, 2, 0, 6, 8, 0, 6, 0} // [1,5,2,0,0,6,8,0,0]
	duplicateZeros(arr)
	fmt.Println(arr)
}

// leetcode submit region begin(Prohibit modification and deletion)
func duplicateZeros(arr []int) {
	// 个人
	cnt, i, n := 0, 0, len(arr)
	for ; i < n-cnt; i++ {
		if arr[i] == 0 {
			cnt++
		}
	}
	i-- // i 必大于 0
	j := n - 1
	if arr[i] == 0 && i+cnt == n { // 末尾是 0，且不复制
		i--
		arr[j] = 0
		j--
		cnt--
	}
	for cnt > 0 {
		if arr[i] == 0 {
			arr[j], arr[j-1] = 0, 0
			j--
			cnt--
		} else {
			arr[j] = arr[i]
		}
		i--
		j--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
