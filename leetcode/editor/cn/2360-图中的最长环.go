package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCycle(edges []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func longestCycle_(edges []int) int {
	// lc
	ans := -1
	visTime := make([]int, len(edges))
	time := 1
	for i := range edges {
		startTime := time
		for i != -1 && visTime[i] == 0 {
			visTime[i], i = time, edges[i]
			time++
		}
		if i != -1 && visTime[i] >= startTime {
			ans = max(ans, time-visTime[i])
		}
	}
	return ans

	// 个人
	//from := make([]int, len(edges)) // 标记一个“集合”，参考并查集的原理
	//ans := 0
	//for f, t := range edges {
	//	if t < 0 { // -1 或者 已访问过
	//		continue
	//	}
	//	tag, i := -2, f // tag 用于计算环的长度，i 用于记录上一个节点
	//	edges[f] = tag  // 将已访问的节点标记为 tag，一方面标记已访问，另一方面用于计算环的长度
	//	for tag--; t >= 0; tag-- {
	//		from[i] = f                       // 标记从 f 节点开始的路径
	//		t, edges[t], i = edges[t], tag, t // 访问下一个节点 && 同 edges[f] = tag && 记录上一个节点
	//	}
	//	if t <= -2 && from[i] == f { // 已访问过，且属于同一个“集合”
	//		ans = max(ans, t-tag)
	//	}
	//}
	//return ans - 1
}
