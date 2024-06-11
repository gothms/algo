package main

import (
	"fmt"
	"math"
)

func main() {
	num := "1101111"
	num = "123456579"                                                                                  // [123,456,579]
	num = "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511" // []
	num = "121202436"                                                                                  // []
	num = "121474836462147483647"                                                                      // [1,2147483646,2147483647]
	fibonacci := splitIntoFibonacci(num)
	fmt.Println(fibonacci)
}

// leetcode submit region begin(Prohibit modification and deletion)
func splitIntoFibonacci(num string) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func splitIntoFibonacci_(num string) []int {
	n := len(num)
	path := make([]int, 0)
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == n {
			return len(path) >= 3 // 长度至少为 3
		}
		t := 0
		if len(path) >= 2 { // 先计算 target
			t = path[len(path)-1] + path[len(path)-2]
			if t > math.MaxInt32 { // 2147483647
				return false
			}
		}
		if num[i] == '0' { // 首字符为 '0'
			if len(path) < 2 || t == 0 {
				path = append(path, 0)
				if dfs(i + 1) {
					return true
				}
				path = path[:len(path)-1]
			}
			return false
		}
		for j, v := i, 0; j < n; j++ {
			v = v*10 + int(num[j]-'0')
			if len(path) < 2 { // 长度不足 3
				if v > math.MaxInt32 {
					break
				} else {
					path = append(path, v)
					if dfs(j + 1) {
						return true
					}
					path = path[:len(path)-1]
				}
			} else { // 斐波那契“序列”
				if v > t {
					break
				} else if v == t {
					path = append(path, v)
					if dfs(j + 1) {
						return true
					}
					path = path[:len(path)-1] // 回溯
				}
			}
		}
		return false
	}
	if dfs(0) && len(path) >= 3 {
		return path
	}
	return nil
}
