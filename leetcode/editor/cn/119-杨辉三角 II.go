package main

import "fmt"

func main() {
	rowIndex := 4
	row := getRow(rowIndex)
	fmt.Println(row)
}

// leetcode submit region begin(Prohibit modification and deletion)
var arr219 [34][]int

func init() {
	for i := 0; i < 34; i++ {
		arr219[i] = make([]int, i+1)
		arr219[i][0], arr219[i][i] = 1, 1
		for j := 1; j < i; j++ {
			arr219[i][j] = arr219[i-1][j-1] + arr219[i-1][j]
		}
	}
}

func getRow(rowIndex int) []int {
	return arr219[rowIndex]

	//// dp
	//dp := make([]int, rowIndex+1)
	//dp[0] = 1
	//for i := 1; i <= rowIndex; i++ {
	//	dp[i] = 1                     // 处理尾巴
	//	for j := i >> 1; j > 0; j-- { // 倒序
	//		dp[j] += dp[j-1] // 中点
	//		dp[i-j] = dp[j]  // 也可能是中点
	//	}
	//}
	//return dp
}

//leetcode submit region end(Prohibit modification and deletion)
