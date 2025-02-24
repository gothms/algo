package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func replaceElements(arr []int) []int {
	for i, maxV := len(arr)-1, -1; i >= 0; i-- {
		arr[i], maxV = maxV, max(maxV, arr[i])
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)
