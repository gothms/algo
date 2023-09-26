package tree

import (
	"strconv"
	"strings"
)

/*
lc-449：序列化和反序列化二叉搜索树
*/

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
	cache, n := make([]int, 0, len(data)>>1), len(data) // val 是一位数字时，需要的容量最大
	for i, j := 0, 0; i < n; i = j + 1 {                // 先映射节点，避免字符串操作
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
