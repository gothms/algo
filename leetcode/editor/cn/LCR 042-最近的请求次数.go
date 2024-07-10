package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type RecentCounter struct {
	p []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.p = append(this.p, t)
	for l := t - 3000; this.p[0] < l; {
		this.p = this.p[1:]
	}
	return len(this.p)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
//leetcode submit region end(Prohibit modification and deletion)
