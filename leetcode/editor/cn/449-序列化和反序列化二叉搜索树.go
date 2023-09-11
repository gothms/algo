//序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
//
// 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序
//列化为最初的二叉搜索树。
//
// 编码的字符串应尽可能紧凑。
//
//
//
// 示例 1：
//
//
//输入：root = [2,1,3]
//输出：[2,1,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数范围是 [0, 10⁴]
// 0 <= Node.val <= 10⁴
// 题目数据 保证 输入的树是一棵二叉搜索树。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 设计 二叉搜索树 字符串 二叉树 👍 425 👎 0

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// [5,3,1,2,4,8,6,7,10,9]
	// [5,3,8,1,4,6,10,0,2]	// 测试样例是层序，且中间没 nil
	// [50,30,10,20,40,80,60,70,100,90]
	tree := make([]*TreeNode, 1, 11)
	for i := 1; i <= 10; i++ {
		tree = append(tree, &TreeNode{Val: i * 10})
	}
	tree[5].Left, tree[5].Right = tree[3], tree[8]
	tree[3].Left, tree[3].Right = tree[1], tree[4]
	tree[1].Left, tree[1].Right = nil, tree[2]
	tree[8].Left, tree[8].Right = tree[6], tree[10]
	tree[6].Left, tree[6].Right = nil, tree[7]
	tree[10].Left, tree[10].Right = tree[9], nil
	cc := Codec{}
	serialize := cc.serialize(tree[5])
	fmt.Println(serialize)
	deserialize := cc.deserialize(serialize)
	fmt.Println(deserialize)
}

/*
要点
	二叉树遍历为单线程
	后遍历的分支，可能用到先遍历的分支改变了的数据和状态

思路：前序遍历
	1.只是序列化的话，前中后序/层序遍历，都可以，所以关键是考虑反序列化
		也就是通过序列化了的字符串构造出二叉搜索树
	2.根据前后序遍历后的序列，其实反序列化都还简单
		2.1.根节点
			比如前序遍历，那么序列中第一个“节点”就是根节点
		2.2.左右节点
			二叉搜索树的根节点与左右节点存在：left.Val < root.Val < right.Val
			所以记录每个“根”节点的左/右区间，然后就能判断一个子节点是否为这个“根”的左/右节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
// 版本三

type Codec struct{}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		sb.WriteString(strconv.Itoa(r.Val))
		sb.WriteByte(' ') // 最后一个 ' ' 是哨兵
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	cache, n := make([]int, 0, len(data)>>1), len(data)
	for i, j := 0, 0; i < n; i = j + 1 { // 先映射节点，避免字符串操作
		j = i
		for data[j] != ' ' {
			j++
		}
		v, _ := strconv.Atoi(data[i:j]) // 求节点值
		cache = append(cache, v)        // 将字符串映射为“节点”
	}
	l, m := 0, len(cache)
	var build func(int, int) *TreeNode
	build = func(min, max int) *TreeNode {
		if l == m {
			return nil
		}
		val := cache[l]             // 当前判断的“节点”
		if val < min || val > max { // 不属于此区间的节点
			return nil
		}
		l++ // 下一个节点的开始位置
		return &TreeNode{val, build(min, val), build(val, max)}
	}
	return build(-1, 10_001) // 左区间 -1，右区间 10_001
}

// 版本二
//type Codec struct{}
//
//func Constructor() Codec { return Codec{} }
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	var sb strings.Builder
//	var dfs func(*TreeNode)
//	dfs = func(r *TreeNode) {
//		if r == nil {
//			return
//		}
//		sb.WriteString(strconv.Itoa(r.Val)) // 前序遍历
//		sb.WriteByte(' ')                   // 最后一个 ' ' 是哨兵
//		dfs(r.Left)
//		dfs(r.Right)
//	}
//	dfs(root)
//	return sb.String()
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	l, n := 0, len(data)
//	var build func(int, int) *TreeNode
//	build = func(min, max int) *TreeNode {
//		if l == n {
//			return nil
//		}
//		r := l
//		for unicode.IsDigit(rune(data[r])) { // 有可能会查询多次（尤其回溯时）
//			r++
//		}
//		val, _ := strconv.Atoi(data[l:r]) // 节点的 val
//		fmt.Println(l, r, min, val, max)  // 有可能会查询多次
//		if val < min || val > max {       // 不属于此区间的节点
//			return nil
//		}
//		l = r + 1 // 下一个节点的开始位置
//		return &TreeNode{val, build(min, val), build(val, max)}
//	}
//	return build(-1, 10_001) // 左区间 -1，右区间 10_001
//}

// 版本一
//func (this *Codec) serialize(root *TreeNode) string {
//	if root == nil {
//		return ""
//	}
//	var sb strings.Builder
//	var dfs func(*TreeNode)
//	dfs = func(r *TreeNode) {
//		if r == nil {
//			return
//		}
//		dfs(r.Right)
//		dfs(r.Left)
//		sb.WriteString(strconv.Itoa(r.Val))
//		sb.WriteByte(' ')
//	}
//	dfs(root)
//	return sb.String()[:sb.Len()-1]
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	if data == "" {
//		return nil
//	}
//	values := strings.Split(data, " ")
//	last := len(values) - 1
//	fmt.Println(values)
//	var build func(int, int) *TreeNode
//	build = func(min, max int) *TreeNode {
//		if last == -1 {
//			return nil
//		}
//		val, _ := strconv.Atoi(values[last])
//		if val < min || val > max {
//			return nil
//		}
//		last--
//		return &TreeNode{val, build(min, val), build(val, max)}
//	}
//	return build(-1, 10_001)
//}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//leetcode submit region end(Prohibit modification and deletion)
