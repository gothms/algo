package main

import "fmt"

func main() {
	s := "ABC"     // 10
	s = "ABA"      // 8
	s = "LEETCODE" // 92
	s = "LEETCOD"  // 64
	s = "AAA"      // 3
	s = "CABA"     // 16
	s = "CBABA"    // 21
	s = "CBABAB"   // 26
	letterString := uniqueLetterString(s)
	fmt.Println(letterString)
}

// leetcode submit region begin(Prohibit modification and deletion)
func uniqueLetterString(s string) int {
	memo := [26][]int{}
	for i, c := range s {
		j := c - 'A'
		memo[j] = append(memo[j], i)
	}
	ans, n := 0, len(s)
	for _, v := range memo {
		vs := make([]int, 0, len(v)+2)
		vs = append(append(append(vs, -1), v...), n)
		for i := 1; i <= len(v); i++ {
			ans += (vs[i] - vs[i-1]) * (vs[i+1] - vs[i])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func uniqueLetterString_(s string) int {
	// 示例：BCADEAFGA，考虑最后一个 A
	// pre + 8-5 - (5-2)
	last0, last1 := [26]int{}, [26]int{}
	for i := range last0 {
		last0[i], last1[i] = -1, -1
	}
	ans, pre := 0, 0 // pre 记录以上一个字符为结尾的子字符串的 uniqueLetterString() 值
	for i, c := range s {
		j := c - 'A'
		if last0[j] >= 0 {
			pre += i - last0[j] - (last0[j] - last1[j])
			last0[j], last1[j] = i, last0[j]
		} else {
			pre += i + 1
			last0[j] = i
		}
		ans += pre
	}
	return ans

	// dp
	//m := make(map[rune][]int)
	//for i, c := range s {
	//	m[c] = append(m[c], i)
	//}
	//ret, n := 0, len(s)
	//for _, arr := range m {
	//	temp := append(append([]int{-1}, arr...), n)
	//	for i := 1; i <= len(arr); i++ {
	//		ret += (temp[i] - temp[i-1]) * (temp[i+1] - temp[i])
	//	}
	//}
	//return ret
}
