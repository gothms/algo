package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLongestWord(s string, dictionary []string) string {
	n := len(s)
	var ans string
	for _, str := range dictionary {
		m := len(str)
		if m < len(ans) || m == len(ans) && str >= ans {
			continue
		}
		i, j := 0, 0
		for i < n && j < m {
			if s[i] != str[j] {
				i++
				continue
			}
			i++
			j++
		}
		if j == m {
			ans = str
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
