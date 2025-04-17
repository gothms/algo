package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func executeInstructions(n int, startPos []int, s string) []int {
	// 模拟
	m := len(s)
	ans := make([]int, m)
	for i := range s {
		cur := map[rune]int{
			'L': startPos[1],
			'R': n - startPos[1] - 1,
			'U': startPos[0],
			'D': n - startPos[0] - 1}
		ans[i] = m - i
		for j, d := range s[i:] {
			if cur[d] == 0 {
				ans[i] = j
				break
			}
			cur[d]--
			var rev rune
			switch d {
			case 'L':
				rev = 'R'
			case 'R':
				rev = 'L'
			case 'U':
				rev = 'D'
			case 'D':
				rev = 'U'
			}
			cur[rev]++
		}
	}
	return ans

	//cur := map[byte]int{
	//	'L': startPos[0],
	//	'R': n - startPos[0] - 1,
	//	'U': startPos[1],
	//	'D': n - startPos[1] - 1}
	//m := len(s)
	//ans := make([]int, m+1)
	//for i := m - 1; i >= 0; i-- {
	//	d := s[i]
	//	if d == 0 {
	//		continue
	//	}
	//	cur[d]--
	//	var rev byte
	//	switch d {
	//	case 'L':
	//		rev = 'R'
	//	case 'R':
	//		rev = 'L'
	//	case 'U':
	//		rev = 'D'
	//	case 'D':
	//		rev = 'U'
	//	}
	//	cur[rev]++
	//	ans[i]=ans[i+1]+1
	//	for ; j < m && s[j] == rev; {
	//		j=
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
