package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	// 有序集合
	//rbt := redblacktree.NewWithIntComparator()
	//for i, v := range nums {
	//	if ceil, ok := rbt.Ceiling(v - valueDiff); ok && ceil.Key.(int) <= v+valueDiff {
	//		return true
	//	}
	//	rbt.Put(v, struct{}{})
	//	if i >= indexDiff {
	//		rbt.Remove(nums[i-indexDiff])
	//	}
	//}
	//return false

	// 桶
	getIdx := func(v, w int) int {
		if v >= 0 {
			return v / w
		}
		return (v+1)/w - 1
	}
	memo := make(map[int]int, indexDiff+1)
	for i, v := range nums {
		idx := getIdx(v, valueDiff+1)
		if _, ok := memo[idx]; ok {
			return true
		}
		if l, ok := memo[idx-1]; ok && v-l <= valueDiff {
			return true
		}
		if r, ok := memo[idx+1]; ok && r-v <= valueDiff {
			return true
		}
		memo[idx] = v
		if i >= indexDiff {
			delete(memo, getIdx(nums[i-indexDiff], valueDiff+1))
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
