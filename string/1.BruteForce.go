package string

/*
字符串匹配算法
	1.蛮力匹配算法
		前缀蛮力匹配算法
		后缀蛮力匹配算法
	2.Sunday算法
	3.KMP算法
	4.Boyer-Moore算法
	5.Rabin-Karp算法
多模式串匹配算法
	Trie
	AC 自动机

Go 字符串匹配算法：https://juejin.cn/post/6844903503920365582
*/

// BruteForcePrefix 前缀蛮力匹配算法
func BruteForcePrefix(s, p string) []int {
	n, m := len(s)-len(p), len(p)
	prefix := make([]int, 0)
	for i, j := 0, 0; i <= n; i++ {
		for j = 0; j < m && s[i+j] == p[j]; {
			j++
		}
		if j == m {
			prefix = append(prefix, i)
		}
	}
	return prefix
}

// BruteForceSuffix 后缀蛮力匹配算法
func BruteForceSuffix(s, p string) []int {
	n, m := len(s)-len(p), len(p)-1
	suffix := make([]int, 0)
	for i, j := 0, 0; i <= n; i++ {
		for j = m; j >= 0 && s[i+j] == p[j]; {
			j--
		}
		if j == -1 {
			suffix = append(suffix, i)
		}
	}
	return suffix
}
