//你会得到一个字符串 s (索引从 0 开始)，你必须对它执行 k 个替换操作。替换操作以三个长度均为 k 的并行数组给出：indices, sources,
// targets。
//
// 要完成第 i 个替换操作:
//
//
// 检查 子字符串 sources[i] 是否出现在 原字符串 s 的索引 indices[i] 处。
// 如果没有出现， 什么也不做 。
// 如果出现，则用 targets[i] 替换 该子字符串。
//
//
// 例如，如果 s = "abcd" ， indices[i] = 0 , sources[i] = "ab"， targets[i] = "eee" ，那么
//替换的结果将是 "eeecd" 。
//
// 所有替换操作必须 同时 发生，这意味着替换操作不应该影响彼此的索引。测试用例保证元素间不会重叠 。
//
//
// 例如，一个 s = "abc" ， indices = [0,1] ， sources = ["ab"，"bc"] 的测试用例将不会生成，因为 "ab"
//和 "bc" 替换重叠。
//
//
// 在对 s 执行所有替换操作后返回 结果字符串 。
//
// 子字符串 是字符串中连续的字符序列。
//
//
//
// 示例 1：
//
//
//
//
//输入：s = "abcd", indexes = [0,2], sources = ["a","cd"], targets = ["eee","ffff"]
//
//输出："eeebffff"
//解释：
//"a" 从 s 中的索引 0 开始，所以它被替换为 "eee"。
//"cd" 从 s 中的索引 2 开始，所以它被替换为 "ffff"。
//
//
// 示例 2：
//
//
//输入：s = "abcd", indexes = [0,2], sources = ["ab","ec"], targets = ["eee",
//"ffff"]
//输出："eeecd"
//解释：
//"ab" 从 s 中的索引 0 开始，所以它被替换为 "eee"。
//"ec" 没有从原始的 S 中的索引 2 开始，所以它没有被替换。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// k == indices.length == sources.length == targets.length
// 1 <= k <= 100
// 0 <= indexes[i] < s.length
// 1 <= sources[i].length, targets[i].length <= 50
// s 仅由小写英文字母组成
// sources[i] 和 targets[i] 仅由小写英文字母组成
//
//
// Related Topics 数组 字符串 排序 👍 88 👎 0

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "vmokgggqzp" // vbfrssozp
	indices := []int{3, 5, 1}
	sources := []string{"kg", "ggq", "mo"}
	targets := []string{"s", "so", "bfr"}
	s = "mhnbzxkwzxtaanmhtoirxheyanoplbvjrovzudznmetkkxrdmr"
	// mhnbzxkwzxtaanmhtoirxhe aqz noplbvjrovzudznmetkkxrdmr
	// mhnbzxkwzxtaanmhtoirxhe aqz noplbvjrovzudznmetkkxrdmr
	// mhnbzxkwzxtaanmhtoirxhe  ya noplbvjrovzudznmetkkxrdmr 问题在23
	indices = []int{46, 29, 2, 44, 31, 26, 42, 9, 38, 23, 36, 12, 16, 7, 33, 18}
	// 2 7 9 12 16 18 23 26 29 31 33 36 38 42 44 46
	sources = []string{"rym", "kv", "nbzxu", "vx", "js", "tp", "tc", "jta", "zqm", "ya", "uz", "avm", "tz", "wn", "yv", "ird"}
	targets = []string{"gfhc", "uq", "dntkw", "wql", "s", "dmp", "jqi", "fp", "hs", "aqz", "ix", "jag", "n", "l", "y", "zww"}
	replaceString := findReplaceString(s, indices, sources, targets)
	fmt.Println(replaceString)
}

/*
思路：模拟
	1.排序
		按照集合 indices 值的先后顺序来判断是否替换 s，所以先对 indices 的值进行排序
		indices, sources, targets 三个集合的元素是映射好的
		所以可以按 indices 的值进行索引排序
	2.模拟
		排序好后，就可以按题意对 s 进行遍历模拟
		和 sources 中的字符串匹配，就替换，不匹配就保持 s 对应位置的字符
		其他都是一些细节处理
*/
// leetcode submit region begin(Prohibit modification and deletion)
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	// 模拟
	n, m := len(s), len(indices)
	idx := make([]int, m)
	for i := 1; i < m; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { // 排序indices的索引
		return indices[idx[i]] < indices[idx[j]]
	})
	var sb strings.Builder
	for i, j := 0, 0; i < n; i++ {
		if i != indices[idx[j]] { // indices值不匹配
			sb.WriteByte(s[i])
			continue
		}
		ss := sources[idx[j]]
		if end := i + len(ss); end <= n && s[i:end] == ss { // sources匹配，s不能越界
			sb.WriteString(targets[idx[j]])
			i = end - 1
		} else { // sources不匹配
			sb.WriteByte(s[i])
		}
		j++
		if j == m { // indices遍历完了
			sb.WriteString(s[i+1:])
			break
		}
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
