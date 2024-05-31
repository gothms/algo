package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fleet := carFleet(target, position, speed)
	fmt.Println(fleet)
}

// leetcode submit region begin(Prohibit modification and deletion)
func carFleet(target int, position []int, speed []int) int {
	ids := make([]int, len(position))
	for i := range ids {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return position[ids[i]] > position[ids[j]]
	})
	ans, d, s := 0, 0, 1
	for _, i := range ids {
		if (target-position[i])*s > speed[i]*d {
			d, s = target-position[i], speed[i]
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
