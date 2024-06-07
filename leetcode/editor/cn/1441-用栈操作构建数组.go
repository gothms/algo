package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func buildArray(target []int, n int) []string {

}

//leetcode submit region end(Prohibit modification and deletion)

func buildArray_(target []int, n int) []string {
	ans := make([]string, 0)
	cur := 1
	for _, v := range target {
		for ; cur < v; cur++ {
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
		cur++
	}
	return ans
}
