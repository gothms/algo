package main

func main() {

}

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool           { return true }
func (this NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger { return nil }

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	arr []int
	i   int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	// 栈

	// 个人：dfs
	arr := make([]int, 0)
	var dfs func([]*NestedInteger)
	dfs = func(ni []*NestedInteger) {
		for _, list := range ni { // 解析
			if list.IsInteger() { // 数字
				arr = append(arr, list.GetInteger())
			} else {
				dfs(list.GetList()) // 集合
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{arr, 0}
}

func (this *NestedIterator) Next() int {
	v := this.arr[this.i]
	this.i++
	return v
}

func (this *NestedIterator) HasNext() bool {
	return len(this.arr) > this.i
}

//leetcode submit region end(Prohibit modification and deletion)
