//给定一个字符串 s 和一个字符串字典
// wordDict ，在字符串
// s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
//
// 注意：词典中的同一个单词可能在分段中被重复使用多次。
//
//
//
// 示例 1：
//
//
//输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
//输出:["cats and dog","cat sand dog"]
//
//
// 示例 2：
//
//
//输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine",
//"pineapple"]
//输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
//解释: 注意你可以重复使用字典中的单词。
//
//
// 示例 3：
//
//
//输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
//输出:[]
//
//
//
//
// 提示：
//
//
//
//
//
// 1 <= s.length <= 20
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 10
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中所有字符串都 不同
//
//
// Related Topics 字典树 记忆化搜索 数组 哈希表 字符串 动态规划 回溯 👍 722 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	i := wordBreak(s, wordDict)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) []string {
	// 记忆化搜索
	n := len(s)
	ret, path, memo := make([]string, 0), make([]int, 0), make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		memo[w] = true
	}
	var sb strings.Builder
	var dfs func(int)
	dfs = func(i int) {
		if i == n { // 组装字符串
			m := len(path)
			for j := 1; j < m; j++ {
				sb.WriteString(s[path[j-1]:path[j]])
				sb.WriteByte(' ')
			}
			sb.WriteString(s[path[m-1]:n]) // 收尾
			ret = append(ret, sb.String())
			sb.Reset()
			return
		}
		path = append(path, i) // 记录起始点（比终止点更方便）
		for j := i + 1; j <= n; j++ {
			if memo[s[i:j]] { // 可切分
				dfs(j)
			}
		}
		path = path[:len(path)-1] // 回溯
	}
	dfs(0)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
