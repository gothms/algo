package main

import "strconv"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func calPoints(operations []string) int {
	st := make([]int, 0)
	for _, o := range operations {
		switch o {
		case "C":
			st = st[:len(st)-1]
		case "D":
			st = append(st, st[len(st)-1]<<1)
		case "+":
			st = append(st, st[len(st)-1]+st[len(st)-2])
		default:
			v, _ := strconv.Atoi(o)
			st = append(st, v)
		}
	}
	ans := 0
	for _, v := range st {
		ans += v
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
