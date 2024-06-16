package main

func main() {

}

var f5125 []int

func init() {
	const n = 21
	f5125 = make([]int, n)
	f5125[1] = 1
	for i := 2; i < n; i++ {
		f5125[i] = f5125[i-1]<<1 + 2
	}
	//fmt.Println(f5125)
}

func eatPeaches(nDays int) int {
	return f5125[nDays]
}
