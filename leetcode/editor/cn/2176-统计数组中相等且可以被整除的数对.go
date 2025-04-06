package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(nums []int, k int) int {
	ans := 0
	memo := make(map[int][]int)
	for i, v := range nums {
		if idx, ok := memo[v]; ok {
			for _, j := range idx {
				if i*j%k == 0 {
					ans++
				}
			}
		}
		memo[v] = append(memo[v], i)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
