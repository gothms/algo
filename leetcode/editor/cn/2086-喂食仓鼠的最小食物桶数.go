package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumBuckets(hamsters string) int {
	ans, pos, n := 0, -2, len(hamsters)
	for i, c := range hamsters {
		if c == '.' {
			continue
		}
		if pos == i-1 {
		} else if i+1 < n && hamsters[i+1] == '.' {
			pos = i + 1
			ans++
		} else if i > 0 && hamsters[i-1] == '.' {
			ans++
		} else {
			return -1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
