package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func hardestWorker(n int, logs [][]int) int {
	id, last, maxT := 0, 0, 0
	for _, v := range logs {
		if t := v[1] - last; t > maxT || t == maxT && v[0] < id {
			id, maxT = v[0], t
		}
		last = v[1]
	}
	return id
}

//leetcode submit region end(Prohibit modification and deletion)
