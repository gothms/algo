package main

import "fmt"

func main() {
	nums := []int{4, 2, 5, 7}
	nums = []int{4, 1, 2, 1}
	nums = []int{648, 831, 560, 986, 192, 424, 997, 829, 897, 843}
	ii := sortArrayByParityII(nums)
	fmt.Println(ii)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if (nums[i]^i)&1 == 0 {
			continue
		}
		j := i + 1
		for ; j < len(nums); j += 2 {
			if (nums[j]^j)&1 > 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
