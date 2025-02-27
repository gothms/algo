package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type BrowserHistory struct {
	st []string
	i  int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.i++
	this.st = append(this.st[:this.i], url)
}

func (this *BrowserHistory) Back(steps int) string {
	this.i = max(0, this.i-steps)
	return this.st[this.i]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.i = min(len(this.st)-1, this.i+steps)
	return this.st[this.i]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
//leetcode submit region end(Prohibit modification and deletion)
