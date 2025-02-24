package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findSpecialInteger(arr []int) int {
	n := len(arr)
	k := n / 4 // i+k
	for i, v := range arr {
		if j := i + k; j < n && v == arr[j] {
			return v
		}
	}
	return arr[0]
}

//leetcode submit region end(Prohibit modification and deletion)
