package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfWeeks(milestones []int) int64 {
	maxV, sum := 0, 0
	for _, v := range milestones {
		maxV = max(maxV, v)
		sum += v
	}
	if maxV > sum-maxV {
		return int64(sum-maxV)<<1 + 1
	}
	return int64(sum)
}

//leetcode submit region end(Prohibit modification and deletion)
