package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumOperations(num string) int {
	// lc：以 00 25 50 75 结尾，如果仅含一个 0，则 len(num)-1

	// 个人
	//f := [25]int{}
	//for i := range f {
	//	f[i] = -1
	//}
	//f[0] = 0
	//for _, c := range num {
	//	v := int(c - '0')
	//	for i, m := range f {
	//		if m >= 0 {
	//			j := (i*10 + v) % 25
	//			f[j] = max(f[j], m+1)
	//		}
	//	}
	//}
	//return len(num) - f[0]
}

//leetcode submit region end(Prohibit modification and deletion)
