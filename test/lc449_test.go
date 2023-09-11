package test

import (
	"strconv"
	"strings"
	"testing"
)

func Benchmark449FuncDfs(b *testing.B) {
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serialize := cc.serializeFunc(tree[5])
		//fmt.Println(serialize)
		cc.deserializeFunc(serialize)
		//fmt.Println(deserialize)
	}
	b.StopTimer()
}
func Benchmark449DFS(b *testing.B) {
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		serialize := cc.serialize(tree[5])
		//fmt.Println(serialize)
		cc.deserialize(serialize)
		//fmt.Println(deserialize)
	}
	b.StopTimer()
}

/*
Benchmark449FuncDfs-8            1640522               740.5 ns/op           432 B/op         15 allocs/op
Benchmark449DFS-8                1653282               737.3 ns/op           432 B/op         15 allocs/op
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 版本三
type Codec struct{}

func Constructor() Codec { return Codec{} }
func (this *Codec) serializeFunc(root *TreeNode) string {
	var sb strings.Builder
	dfs449(root, &sb)
	return sb.String()
}
func dfs449(r *TreeNode, sb *strings.Builder) {
	if r == nil {
		return
	}
	sb.WriteString(strconv.Itoa(r.Val)) // 前序遍历
	sb.WriteByte(' ')                   // 最后一个 ' ' 是哨兵
	dfs449(r.Left, sb)
	dfs449(r.Right, sb)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeFunc(data string) *TreeNode {
	cache, n, l := make([]int, 0, len(data)>>1), len(data), 0
	for i, j := 0, 0; i < n; i = j + 1 {
		j = i
		for data[j] != ' ' {
			j++
		}
		v, _ := strconv.Atoi(data[i:j])
		cache = append(cache, v)
	}
	return dfsDe449(-1, 10_001, &l, cache) // 左区间 -1，右区间 10_001
}
func dfsDe449(min, max int, i *int, arr []int) *TreeNode {
	if *i == len(arr) {
		return nil
	}
	val := arr[*i]
	if val < min || val > max { // 不属于此区间的节点
		return nil
	}
	*i++ // 下一个节点的开始位置
	return &TreeNode{val, dfsDe449(min, val, i, arr), dfsDe449(val, max, i, arr)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		sb.WriteString(strconv.Itoa(r.Val)) // 前序遍历
		sb.WriteByte(' ')                   // 最后一个 ' ' 是哨兵
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	cache, n := make([]int, 0, len(data)>>1), len(data)
	for i, j := 0, 0; i < n; i = j + 1 {
		j = i
		for data[j] != ' ' {
			j++
		}
		v, _ := strconv.Atoi(data[i:j])
		cache = append(cache, v)
	}
	l, m := 0, len(cache)
	var build func(int, int) *TreeNode
	build = func(min, max int) *TreeNode {
		if l == m {
			return nil
		}
		val := cache[l]
		if val < min || val > max { // 不属于此区间的节点
			return nil
		}
		l++ // 下一个节点的开始位置
		return &TreeNode{val, build(min, val), build(val, max)}
	}
	return build(-1, 10_001) // 左区间 -1，右区间 10_001
}
