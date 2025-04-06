package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
var memo1925 map[int]struct{}

func init() {
	n := 250
	memo1925 = make(map[int]struct{}, n)
	for i := 1; i <= n; i++ {
		memo1925[i*i] = struct{}{}
	}
}
func countTriples(n int) int {
	ans, mx := 0, n*n
	for i := 3; i < n; i++ {
		for j := 3; j < i; j++ {
			k := i*i + j*j
			if k > mx {
				break
			}
			if _, ok := memo1925[k]; ok {
				ans += 2
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
