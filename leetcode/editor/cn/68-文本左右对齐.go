//给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
// 你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
//
// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
// 注意:
//
//
// 单词是指由非空格字符组成的字符序列。
// 每个单词的长度大于 0，小于等于 maxWidth。
// 输入单词数组 words 至少包含一个单词。
//
//
//
//
// 示例 1:
//
//
//输入: words = ["This", "is", "an", "example", "of", "text", "justification."],
//maxWidth = 16
//输出:
//[
//   "This    is    an",
//   "example  of text",
//   "justification.  "
//]
//
//
// 示例 2:
//
//
//输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
//输出:
//[
//  "What   must   be",
//  "acknowledgment  ",
//  "shall be        "
//]
//解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//     因为最后一行应为左对齐，而不是左右两端对齐。
//     第二行同样为左对齐，这是因为这行只包含一个单词。
//
//
// 示例 3:
//
//
//输入:words = ["Science","is","what","we","understand","well","enough","to",
//"explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 2
//0
//输出:
//[
//  "Science  is  what we",
//  "understand      well",
//  "enough to explain to",
//  "a  computer.  Art is",
//  "everything  else  we",
//  "do                  "
//]
//
//
//
//
// 提示:
//
//
// 1 <= words.length <= 300
// 1 <= words[i].length <= 20
// words[i] 由小写英文字母和符号组成
// 1 <= maxWidth <= 100
// words[i].length <= maxWidth
//
//
// Related Topics 数组 字符串 模拟 👍 372 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	justify := fullJustify(words, maxWidth)
	for i, s := range justify {
		fmt.Println(i, s)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func fullJustify(words []string, maxWidth int) []string {
	const Space = ' '
	cnt, j, n := len(words[0]), 0, len(words)
	ret := make([]string, 0)
	var sb strings.Builder
	sBuild := func(i, j int, cnt int) {
		var sc, c int
		start := j // 标记起始点
		if i-j == 1 {
			sc = maxWidth - cnt // 只有一个单词
		} else {
			sc, c = (maxWidth-cnt)/(i-j-1)+1, (maxWidth-cnt)%(i-j-1) // 前 c 个间隔为 sc+1 个空格
		}
		for k := 0; j < i; j++ {
			sb.WriteString(words[j])
			if i-start > 1 && i-j == 1 { // 单词数 > 1，且添加最后一个单词
				break
			}
			for m := 0; m < sc; m++ {
				sb.WriteByte(Space)
			}
			if k < c { // 前 c 个间隔为 sc+1 个空格
				sb.WriteByte(Space)
				k++
			}
		}
	}
	for i := 1; i < n; i++ {
		if cnt+len(words[i])+1 > maxWidth { // 延迟
			sBuild(i, j, cnt)
			ret = append(ret, sb.String()) // 当前行
			sb.Reset()
			cnt = len(words[i]) // 下一行
			j = i
		} else {
			cnt += len(words[i]) + 1 // 加上一个空格
		}
	}
	sBuild(n, j, maxWidth)
	for ; cnt < maxWidth; cnt++ { // 最后一行要求是在末尾补全空格
		sb.WriteByte(Space)
	}
	ret = append(ret, sb.String())
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
