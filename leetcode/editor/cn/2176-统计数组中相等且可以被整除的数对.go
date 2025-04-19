package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(nums []int, k int) int {
	memo := make(map[int][]int)
	ans := 0
	for i, v := range nums {
		if arr := memo[v]; len(arr) > 0 {
			for _, j := range arr {
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
