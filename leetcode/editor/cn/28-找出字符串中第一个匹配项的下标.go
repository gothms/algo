package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	// RabinKarp

	// kmp
	kmp := func(p string) []int {
		n := len(p)
		prefix := make([]int, n)
		k := -1
		prefix[0] = k
		for i := 1; i < n; i++ {
			for k != -1 && p[k+1] != p[i] {
				k = prefix[k]
			}
			if p[k+1] == p[i] {
				k++
			}
			prefix[i] = k
		}
		return prefix
	}
	prefix := kmp(needle)
	m, n := len(haystack), len(needle)
	for i, j := 0, 0; i < m; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = prefix[j-1] + 1
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
