package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countOperations(num1 int, num2 int) int {
	// lc：辗转相除法
	ans := 0
	for num1 > 0 {
		ans += num2 / num1
		num1, num2 = num2%num1, num1
	}
	return ans

	// 个人
	//if num1 == 0 || num2 == 0 {
	//	return 0
	//}
	//ans := 0
	//for num1 != num2 {
	//	var d int
	//	if num1 > num2 {
	//		d = (num1-num2-1)/num2 + 1
	//		num1 -= d * num2
	//	} else {
	//		d = (num2-num1-1)/num1 + 1
	//		num2 -= d * num1
	//	}
	//	ans += d
	//}
	//if num1 > 0 {
	//	ans++
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
