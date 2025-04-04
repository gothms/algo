package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	uni := [26]int{}
	for _, c := range s {
		uni[c-'a']++
	}
	for _, c := range t {
		uni[c-'a']--
	}
	for _, v := range uni {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
