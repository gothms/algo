package main

import "fmt"

func main() {
	distance := []int{2, 1, 1, 2}      // true
	distance = []int{1, 2, 3, 4}       // false
	distance = []int{1, 1, 1, 2, 1}    // true
	distance = []int{3, 3, 4, 2, 2}    // false
	distance = []int{3, 3, 3, 2, 1, 1} // false
	crossing := isSelfCrossing(distance)
	fmt.Println(crossing)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isSelfCrossing(distance []int) bool {
	i, n := 0, len(distance)
	for i < n && (i < 2 || distance[i] > distance[i-2]) { // 外向螺旋
		i++
	}
	if i == n {
		return false
	}
	// 变形
	if i == 3 && distance[i] == distance[i-2] || // 比如 i == 3 distance = [1,1,3,1,...]
		i >= 4 && distance[i] >= distance[i-2]-distance[i-4] { // 比如 i == 4 distance = [2,1,3,2,2,...]
		distance[i-1] -= distance[i-3] // 修正：只能开始内向螺旋
	}
	i++                                        // 成功变形
	for i < n && distance[i] < distance[i-2] { // 内向螺旋
		i++
	}
	return i != n
}

//leetcode submit region end(Prohibit modification and deletion)
