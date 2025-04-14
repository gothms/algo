package algo

/*
负整数 转 uint
	源码
		num := -3
		cur := bits.TrailingZeros(uint(num)) // 正常
		cur = bits.TrailingZeros(uint(-3))   // 报错
		fmt.Println(cur)
	报错
		# command-line-arguments
		.\面试题 05.03-翻转数位.go:16:32: constant -3 overflows uint
切片
	arr := []int{0}
	arr[0], arr[0] = arr[0]+1, arr[0]-1
	fmt.Println(arr) // -1
*/
