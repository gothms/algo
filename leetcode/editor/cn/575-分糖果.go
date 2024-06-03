package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(candyType []int) int {
	memo := make(map[int]struct{})
	for _, v := range candyType {
		memo[v] = struct{}{}
	}
	return min(len(memo), len(candyType)>>1)
}

//leetcode submit region end(Prohibit modification and deletion)
