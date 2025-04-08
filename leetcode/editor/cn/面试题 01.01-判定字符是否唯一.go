package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isUnique(astr string) bool {
	memo := [26]bool{}
	for _, c := range astr {
		i := int(c - 'a')
		if memo[i] {
			return false
		}
		memo[i] = true
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
