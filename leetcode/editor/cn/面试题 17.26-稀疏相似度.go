//两个(具有不同单词的)文档的交集(intersection)中元素的个数除以并集(union)中元素的个数，就是这两个文档的相似度。例如，{1, 5, 3}
// 和 {1, 7, 2, 3} 的相似度是 0.4，其中，交集的元素有 2 个，并集的元素有 5 个。给定一系列的长篇文档，每个文档元素各不相同，并与一个
//ID 相关联。它们的相似度非常“稀疏”，也就是说任选 2 个文档，相似度都很接近 0。请设计一个算法返回每对文档的 ID 及其相似度。只需输出相似度大于 0 的组合
//。请忽略空文档。为简单起见，可以假定每个文档由一个含有不同整数的数组表示。
//
// 输入为一个二维数组 docs，docs[i] 表示 id 为 i 的文档。返回一个数组，其中每个元素是一个字符串，代表每对相似度大于 0 的文档，其格式为
// {id1},{id2}: {similarity}，其中 id1 为两个文档中较小的 id，similarity 为相似度，精确到小数点后 4 位。以任意顺序
//返回数组均可。
//
// 示例:
//
// 输入:
//[
//  [14, 15, 100, 9, 3],
//  [32, 1, 9, 3, 5],
//  [15, 29, 2, 6, 8, 7],
//  [7, 10]
//]
//输出:
//[
//  "0,1: 0.2500",
//  "0,2: 0.1000",
//  "2,3: 0.1429"
//]
//
// 提示：
//
//
// docs.length <= 500
// docs[i].length <= 500
//
//
// Related Topics 数组 哈希表 排序 👍 43 👎 0

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	f := 3.1415926
	fv := math.Trunc(f*1e2 + 0.5)
	fmt.Println(fv)
	fv = math.Trunc(f*1e2+0.5) * 1e-2
	fmt.Println(fv)
	// Sprintf 保留2位，再转为 float64
	fv, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
	fmt.Println(fv)

	docs := [][]int{{14, 15, 100, 9, 3},
		{32, 1, 9, 3, 5},
		{15, 29, 2, 6, 8, 7},
		{7, 10}}
	similarities := computeSimilarities(docs)
	fmt.Println(similarities)
}

// leetcode submit region begin(Prohibit modification and deletion)
func computeSimilarities(docs [][]int) []string {
	// 原理：求 n 个集合之间的交集元素个数：https://leetcode.cn/problems/sparse-similarity-lcci/solutions/298148/c-yuan-shu-jie-fa-shi-xian-ha-xi-biao-by-wen-zhong/
	n := len(docs)
	d2d, mp := make([][]int, n-1), make(map[int][]int)
	for i, doc := range docs {
		for _, d := range doc {
			mp[d] = append(mp[d], i) // 元素：出现的集合
		}
	}
	//fmt.Println(mp)
	for i := 0; i < n-1; i++ {
		d2d[i] = make([]int, n)
	}
	for _, v := range mp {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				d2d[v[i]][v[j]]++ // 统计交集数
			}
		}
	}
	//fmt.Println(d2d)
	ret := make([]string, 0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if d := d2d[i][j]; d > 0 { // 集合 i 和 j 之间，交集的个数
				v := float64(d) / float64(len(docs[i])+len(docs[j])-d)
				ret = append(ret, strconv.Itoa(i)+","+strconv.Itoa(j)+": "+fmt.Sprintf("%.4f", v+1e-9))
			} // 格式化修正：+1e-9
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
