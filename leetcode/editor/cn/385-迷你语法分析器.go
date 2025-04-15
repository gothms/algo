package main

import "unicode"

func main() {

}

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool           {}
func (n NestedInteger) GetInteger() int           {}
func (n *NestedInteger) SetInteger(value int)     {}
func (n *NestedInteger) Add(elem NestedInteger)   {}
func (n NestedInteger) GetList() []*NestedInteger {}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	idx, n := 0, len(s)
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}
		if s[idx] == '[' {
			idx++
			for s[idx] != ']' {
				ni.Add(*dfs())
				if s[idx] == ',' {
					idx++
				}
			}
			idx++
			return ni
		}

		negative := s[idx] == '-'
		if negative {
			idx++
		}
		v := 0
		for ; idx < n && unicode.IsNumber(rune(s[idx])); idx++ {
			v = v*10 + int(s[idx]-'0')
		}
		if negative {
			v = -v
		}
		ni.SetInteger(v)
		return ni
	}
	return dfs()
}

//leetcode submit region end(Prohibit modification and deletion)
