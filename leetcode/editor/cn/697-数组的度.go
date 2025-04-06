package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findShortestSubArray(nums []int) int {
	memo := make(map[int][3]int) // startIdx len cnt
	for i, v := range nums {
		if arr, ok := memo[v]; ok {
			arr[1] = i - arr[0] + 1
			arr[2]++
			memo[v] = arr
		} else {
			memo[v] = [3]int{i, 1, 1}
		}
	}
	ans, mx := len(nums), 0
	for _, arr := range memo {
		if arr[2] > mx || arr[2] == mx && arr[1] < ans {
			ans, mx = arr[1], arr[2]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
