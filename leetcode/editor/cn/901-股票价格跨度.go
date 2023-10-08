//设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。
//
// 当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
//
//
// 例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。
//
//
// 实现 StockSpanner 类：
//
//
// StockSpanner() 初始化类对象。
// int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。
//
//
//
//
// 示例：
//
//
//输入：
//["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
//[[], [100], [80], [60], [70], [60], [75], [85]]
//输出：
//[null, 1, 1, 1, 2, 1, 4, 6]
//
//解释：
//StockSpanner stockSpanner = new StockSpanner();
//stockSpanner.next(100); // 返回 1
//stockSpanner.next(80);  // 返回 1
//stockSpanner.next(60);  // 返回 1
//stockSpanner.next(70);  // 返回 2
//stockSpanner.next(60);  // 返回 1
//stockSpanner.next(75);  // 返回 4 ，因为截至今天的最后 4 个股价 (包括今天的股价 75) 都小于或等于今天的股价。
//stockSpanner.next(85);  // 返回 6
//
//
//
//
// 提示：
//
//
// 1 <= price <= 10⁵
// 最多调用 next 方法 10⁴ 次
//
//
// Related Topics 栈 设计 数据流 单调栈 👍 354 👎 0

package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type StockSpanner struct {
	stack [][2]int // 单调递减
	n     int      // 索引
}

func Constructor() StockSpanner {
	return StockSpanner{stack: [][2]int{{math.MaxInt32, -1}}} // 哨兵
}
func (this *StockSpanner) Next(price int) int {
	last := len(this.stack) - 1
	for this.stack[last][0] <= price { // 单调性
		last--
	}
	this.stack = this.stack[:last+1]
	this.stack = append(this.stack, [2]int{price, this.n}) // 添加
	this.n++
	return this.n - this.stack[last][1] - 1 // 计算
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
//leetcode submit region end(Prohibit modification and deletion)
