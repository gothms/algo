package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func largestVariance(s string) int {
	// dp

}

//leetcode submit region end(Prohibit modification and deletion)

func largestVariance_(s string) int {
	ans := 0
	for ax := 'a'; ax <= 'z'; ax++ {
		for in := 'a'; in <= 'z'; in++ {
			f0, f1 := 0, math.MinInt
			for _, c := range s {
				switch c {
				case ax:
					f0 = max(f0, 0) + 1
					f1++
				case in:
					f0 = max(f0, 0) - 1
					f1 = f0
				default:
				}
				ans = max(ans, f1)
			}
		}
	}
	return ans

	//pos := make(map[rune][]int)
	//for i, c := range s {
	//	pos[c] = append(pos[c], i)
	//}
	//ans := 0
	//for cMax, pMax := range pos {
	//	for cMin, pMin := range pos {
	//		if cMax == cMin {
	//			continue
	//		}
	//		mxCnt, dis := 0, math.MinInt32
	//		i, j, m, n := 0, 0, len(pMax), len(pMin)
	//		for i < m || j < n {
	//			if j == n || i < m && pMax[i] < pMin[j] {
	//				mxCnt, dis = max(mxCnt, 0)+1, dis+1
	//				i++
	//			} else {
	//				//mxCnt, dis = max(mxCnt, 0)-1, max(mxCnt, dis, 0)-1
	//				mxCnt = max(mxCnt, 0) - 1
	//				dis = mxCnt
	//				j++
	//			}
	//			ans = max(ans, dis)
	//		}
	//	}
	//}
	//return ans
}
