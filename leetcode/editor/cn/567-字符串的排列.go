package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	// 滑动窗体

	// 双指针
	n, m := len(s1), len(s2)
	if m < n {
		return false
	}
	memo := [26]int{}
	for _, c := range s1 {
		memo[c-'a']--
	}
	left := 0
	for right, c := range s2 {
		idx := int(c - 'a')
		memo[idx]++
		for memo[idx] > 0 {
			memo[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
