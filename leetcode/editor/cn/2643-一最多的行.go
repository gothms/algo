package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func rowAndMaximumOnes(mat [][]int) []int {
	idx, cnt := 0, 0
	for i, row := range mat {
		c := 0
		for _, v := range row {
			c += v
		}
		if c > cnt {
			idx, cnt = i, c
		}
	}
	return []int{idx, cnt}
}

//leetcode submit region end(Prohibit modification and deletion)
