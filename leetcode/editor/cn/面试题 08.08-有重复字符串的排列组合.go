package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func permutation(S string) []string {
	ans := make([]string, 0)
	buf := []byte(S)
	n := len(buf)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(buf))
			return
		}
		dfs(i + 1)
	out:
		for j := i + 1; j < n; j++ {
			for k := i; k < j; k++ {
				if buf[k] == buf[j] {
					continue out
				}
			}
			buf[i], buf[j] = buf[j], buf[i]
			dfs(i + 1)
			buf[i], buf[j] = buf[j], buf[i]
		}
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
