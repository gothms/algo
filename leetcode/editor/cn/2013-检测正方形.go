package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type DetectSquares struct {
	m map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{make(map[int]map[int]int)}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if this.m[x] == nil {
		this.m[x] = make(map[int]int)
	}
	this.m[x][y]++
}

func (this *DetectSquares) Count(point []int) int {
	ans := 0
	x1, y1 := point[0], point[1]
	if this.m[x1] == nil {
		return ans
	}
	for y2, c := range this.m[x1] {
		if d := y1 - y2; d != 0 { // 如果 d==0，则不是正方形
			x2, x3 := x1+d, x1-d
			ans += c * (this.m[x2][y2]*this.m[x2][y1] + this.m[x3][y2]*this.m[x3][y1])
		}
	}
	return ans
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
//leetcode submit region end(Prohibit modification and deletion)
