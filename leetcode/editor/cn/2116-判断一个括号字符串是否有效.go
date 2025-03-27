package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func canBeValid(s string, locked string) bool {
	
}

//leetcode submit region end(Prohibit modification and deletion)

func canBeValid_(s string, locked string) bool {
	if len(s)&1 != 0 {
		return false
	}
	mn, mx := 0, 0
	for i, v := range locked {
		switch v {
		case '1':
			d := 1 - int(s[i])&1<<1 // ( +1，) -1
			mx += d
			if mx < 0 {
				return false
			}
			mn += d
		case '0':
			mn--
			mx++
		}
		if mn < 0 {
			mn = 1
		}
	}
	return mn == 0
}
