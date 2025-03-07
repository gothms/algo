package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n == 1 {
		return ""
	}
	i := 0
	for i < n && palindrome[i] == 'a' ||
		n&1 == 1 && i == n>>1 { // 长度为奇数，且 i == n>>1
		i++
	}
	if i < n { // 将 palindrome[i] 换成 a
		return palindrome[:i] + "a" + palindrome[i+1:]
	}
	return palindrome[:n-1] + "b" // 全是 a
}

//leetcode submit region end(Prohibit modification and deletion)
