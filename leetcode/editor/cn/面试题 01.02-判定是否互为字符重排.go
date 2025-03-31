package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	memo := [26]int{}
	for i := range s1 {
		memo[s1[i]-'a']++
		memo[s2[i]-'a']--
	}
	for _, v := range memo {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
