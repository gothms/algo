package main

import "fmt"

func main() {
	nums := []int{4, 8, 2, 8}
	difference := maximumPrimeDifference(nums)
	fmt.Println(difference)
}

// leetcode submit region begin(Prohibit modification and deletion)
var prime3115 []bool

func init() {
	const n = 101
	prime3115 = make([]bool, n)
	prime3115[1] = true
	for i := 2; i < n; i++ {
		if !prime3115[i] {
			for j := i << 1; j < n; j += i {
				prime3115[j] = true
			}
		}
	}
}
func maximumPrimeDifference(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j && prime3115[nums[i]] {
		i++
	}
	for i < j && prime3115[nums[j]] {
		j--
	}
	return j - i
}

//leetcode submit region end(Prohibit modification and deletion)
