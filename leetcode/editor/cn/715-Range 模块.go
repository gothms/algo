//Range模块是跟踪数字范围的模块。设计一个数据结构来跟踪表示为 半开区间 的范围并查询它们。
//
// 半开区间 [left, right) 表示所有 left <= x < right 的实数 x 。
//
// 实现 RangeModule 类:
//
//
// RangeModule() 初始化数据结构的对象。
// void addRange(int left, int right) 添加 半开区间 [left, right)，跟踪该区间中的每个实数。添加与当前跟踪的
//数字部分重叠的区间时，应当添加在区间 [left, right) 中尚未跟踪的任何数字到该区间中。
// boolean queryRange(int left, int right) 只有在当前正在跟踪区间 [left, right) 中的每一个实数时，才返
//回 true ，否则返回 false 。
// void removeRange(int left, int right) 停止跟踪 半开区间 [left, right) 中当前正在跟踪的每个实数。
//
//
//
//
// 示例 1：
//
//
//输入
//["RangeModule", "addRange", "removeRange", "queryRange", "queryRange",
//"queryRange"]
//[[], [10, 20], [14, 16], [10, 14], [13, 15], [16, 17]]
//输出
//[null, null, null, true, false, true]
//
//解释
//RangeModule rangeModule = new RangeModule();
//rangeModule.addRange(10, 20);
//rangeModule.removeRange(14, 16);
//rangeModule.queryRange(10, 14); 返回 true （区间 [10, 14) 中的每个数都正在被跟踪）
//rangeModule.queryRange(13, 15); 返回 false（未跟踪区间 [13, 15) 中像 14, 14.03, 14.17 这样
//的数字）
//rangeModule.queryRange(16, 17); 返回 true （尽管执行了删除操作，区间 [16, 17) 中的数字 16 仍然会被跟踪）
//
//
//
//
//
// 提示：
//
//
// 1 <= left < right <= 10⁹
// 在单个测试用例中，对 addRange 、 queryRange 和 removeRange 的调用总数不超过 10⁴ 次
//
//
// Related Topics 设计 线段树 有序集合 👍 269 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type RangeModule struct {
	root *treeNode715
}

const n715 = 1e9

type treeNode715 struct {
	left, right *treeNode715
	exist       bool
	d           int
}

func Constructor() RangeModule { return RangeModule{&treeNode715{}} }

func (this *RangeModule) AddRange(left int, right int) {
	this.update(left, right-1, 1, n715, 1, this.root)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.query(left, right-1, 1, n715, this.root)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.update(left, right-1, 1, n715, -1, this.root)
}
func (this *RangeModule) down(cur *treeNode715) {
	if cur.left == nil {
		cur.left = &treeNode715{}
		cur.right = &treeNode715{}
	}
	if cur.d == 0 {
		return
	}
	cur.left.exist = cur.d == 1
	cur.right.exist = cur.d == 1
	cur.left.d = cur.d // 重点：子节点随父节点
	cur.right.d = cur.d
	cur.d = 0
}
func (this *RangeModule) update(f, t, l, r, d int, cur *treeNode715) {
	if f <= l && r <= t {
		cur.exist = d == 1
		cur.d = d // 懒惰标记：唯有区间匹配了，才赋值 d
		return
	}
	this.down(cur) // 否则左右节点为 nil
	m := (l + r) >> 1
	if f <= m {
		this.update(f, t, l, m, d, cur.left)
	}
	if t > m {
		this.update(f, t, m+1, r, d, cur.right)
	}
	cur.exist = cur.left.exist && cur.right.exist
}

func (this *RangeModule) query(f, t, l, r int, cur *treeNode715) bool {
	if f <= l && r <= t {
		return cur.exist
	}
	this.down(cur) // 否则左右节点为 nil
	m := (l + r) >> 1
	var ret = true // 3 个 true
	if f <= m {
		ret = this.query(f, t, l, m, cur.left)
	}
	if t > m {
		ret = ret && this.query(f, t, m+1, r, cur.right)
	}
	//cur.v = cur.left.v + cur.right.v
	return ret
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
