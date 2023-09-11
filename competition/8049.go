package main

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	absVal := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	x, y := absVal(sx-fx), absVal(sy-fy)
	if x+y == 0 {
		return t != 1
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	return maxVal(x, y) <= t
}
