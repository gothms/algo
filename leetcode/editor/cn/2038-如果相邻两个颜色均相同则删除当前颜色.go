package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func winnerOfGame(colors string) bool {
	var (
		a, b int
		last rune
		cnt  int
	)
	for _, c := range colors {
		if c != last {
			last, cnt = c, 1
			continue
		}
		cnt++
		if cnt >= 3 {
			switch c {
			case 'A':
				a++
			case 'B':
				b++
			}
		}
	}
	return a > b
}

//leetcode submit region end(Prohibit modification and deletion)
