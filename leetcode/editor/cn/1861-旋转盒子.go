package main

import "fmt"

func main() {
	box := [][]byte{{'#', '.', '*', '.'},
		{'#', '#', '*', '.'}}
	theBox := rotateTheBox(box)
	fmt.Println(theBox)
}

// leetcode submit region begin(Prohibit modification and deletion)
func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = make([]byte, m)
	}
	for i, b := range box {
		last := n - 1
		for j := last; j >= 0; j-- {
			switch b[j] {
			case '*':
				last = j - 1
			case '#':
				b[last], b[j] = b[j], b[last] // 可能 # 换 #
				last--
				for last > j && b[last] != '.' { // 易错：last > j
					last--
				}
			}
		}
		for j, v := range b {
			ans[j][m-i-1] = v
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
