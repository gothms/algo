package main

import "strconv"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxValue(n string, x int) string {
	m := len(n)
	v := uint8(x + '0')
	i := 0
	if n[0] == '-' {
		i = 1
		for i < m && n[i] <= v {
			i++
		}
	} else {
		for i < m && n[i] >= v {
			i++
		}
	}
	return n[:i] + strconv.Itoa(x) + n[i:]
}

//leetcode submit region end(Prohibit modification and deletion)
