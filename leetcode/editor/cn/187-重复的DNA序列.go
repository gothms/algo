//DNA序列 由一系列核苷酸组成，缩写为
// 'A', 'C', 'G' 和
// 'T'.。
//
//
// 例如，
// "ACGAATTCCG" 是一个 DNA序列 。
//
//
// 在研究 DNA 时，识别 DNA 中的重复序列非常有用。
//
// 给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//输出：["AAAAACCCCC","CCCCCAAAAA"]
//
//
// 示例 2：
//
//
//输入：s = "AAAAAAAAAAAAA"
//输出：["AAAAAAAAAA"]
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 10⁵
// s[i]=='A'、'C'、'G' or 'T'
//
//
// Related Topics 位运算 哈希表 字符串 滑动窗口 哈希函数 滚动哈希 👍 503 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findRepeatedDnaSequences(s string) []string {
	v, ty, n := 0, 1<<18-1, len(s)
	if n <= 10 {
		return nil
	}
	dna2Int := func(c uint8) { // hash 映射：4 进制
		v <<= 2
		switch c {
		case 'A':
		case 'C':
			v++
		case 'G':
			v += 2
		case 'T':
			v += 3
		}
	}
	ret, memo := make([]string, 0), make(map[int]bool)
	for i := 0; i < 10; i++ {
		dna2Int(s[i])
	}
	memo[v] = true
	for i := 10; i < n; i++ {
		v &= ty // 去掉头
		dna2Int(s[i])
		if b, ok := memo[v]; b { // 第二次出现
			ret = append(ret, s[i-9:i+1])
			memo[v] = false
		} else if !ok { // 第一次出现
			memo[v] = true
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
