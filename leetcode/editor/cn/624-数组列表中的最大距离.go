package main

import (
	"fmt"
	"math"
)

func main() {
	arrays := [][]int{{1}, {1}}
	//arrays = [][]int{{1, 5}, {3, 4}}
	distance := maxDistance(arrays)
	fmt.Println(distance)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(arrays [][]int) int {
	minV, lMin, maxV, lMax := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	minId, maxId := -1, -1
	for i, arr := range arrays {
		if arr[0] < minV { // 最小值
			minV, lMin, minId = arr[0], minV, i
		} else { // 第二小的值
			lMin = min(lMin, arr[0])
		}
		if v := arr[len(arr)-1]; v > maxV {
			maxV, lMax, maxId = v, maxV, i
		} else {
			lMax = max(lMax, v)
		}
	}
	if minId == maxId { // 同一个数组
		return max(maxV-lMin, lMax-minV)
	} else {
		return maxV - minV
	}
}

//leetcode submit region end(Prohibit modification and deletion)
