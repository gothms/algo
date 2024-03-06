package main

func main() {
}
func stringCount(n int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	const mod = 1_000_000_007
	if n < 4 {
		return 0
	} else if n == 4 {
		return 12
	}
	cnt := 0
	min := minVal(n, 26)
	v:=6
	for i := 3; i <= min; i++ {
		cnt=(cnt+)%mod
	}
	return cnt
}
