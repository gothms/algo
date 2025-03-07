package main

func main() {

}

/*
思路：散列表
	1.任意 i 散列表记录 memo[target - nums[j]] = j，j区间为 [0,i)
	2.如果 memo[nums[i]] 存在，说明找到 nums[j] + nums[i] = target
*/
// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	for i, v := range nums {
		if j, ok := memo[v]; ok {
			return []int{j, i}
		} else {
			memo[target-v] = i
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
