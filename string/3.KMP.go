package str

import "fmt"

/*
KMP 字符串匹配算法
	前缀匹配算法

前缀表：prefix-table
	最长公共前后缀
	示例：p = "ABABCABAA" prefix = [0 0 1 2 0 1 2 3 1]
		sub 为模式串 p [0,i] 的子串，i<len(p)
		sub 前缀和后缀最长相同元素的长度，记录为 prefix-table
		sub 分别为如下前缀子串时，prefix[i] 的取值：
			A = 0
			AB = 0
			ABA = 1："A" == "A"
			ABAB = 2："AB" == "AB"
			ABABC = 0
			ABABCA = 1："A" == "A"
			ABABCAB = 2："AB" == "AB"
			ABABCABA = 3："ABA" == "ABA"，第一个 "ABA" 的索引为 0 1 2，第二个 "ABA" 的索引为 5 6 7，则 prefix[i] = 3
			ABABCABAA = 1："A" == "A"
KmpLc 写法
	示例：pattern = abababzabababa
		1.提出问题
			最大匹配数：0 0 1 2 3 4 0 1 2 3 4 5 6 ?
			对于字符串 abababzababab
				前缀子串有：a, ab, aba, abab, ababa, ababab, abababz, ...
				后缀子串有：b, ab, bab, abab, babab, ababab, zababab, ...
			所以子字符串 abababzababab 的前缀子串和后缀子串最大匹配了 6 个（ababab），那“次”大匹配了多少呢？
		2.最大匹配数 & 次大匹配数
			容易看出次大匹配了 4 个（abab），更仔细地观察可以发现，次大匹配必定在最大匹配 ababab 中，所以次大匹配数就是 ababab 的最大匹配数
				直接去查算出的表，可以得出该值为 4
			第三大的匹配数同理，它既然比 4 要小，那真前缀和真后缀也只能在 abab 中找，即 abab 的最大匹配数，查表可得该值为 2
				再往下就没有更短的匹配了
		3.问题解答
			计算 ? 的值：
			既然末尾字母不是 z，那么就不能直接 6+1=7 了
			回退到次大匹配 abab，刚好 abab 之后的 a 与末尾的 a 匹配，所以 ? 处的最大匹配数为 5
		总结
			匹配到 ? 时，如何快速定位到 ? 与哪个字符 x 比较？
			核心：
				“最大匹配数”等式：最长可匹配前缀子串 == 最长可匹配后缀子串（简称最长前/后缀）
			转化：
				找到“最长可匹配前缀子串”的最大匹配数，
				则匹配 ? （匹配 “最长后缀”的后缀子串+? 与 “最长前缀”的前缀子串+x）
				即匹配 “最长前缀”的后缀子串+? 与 “最长前缀”的前缀子串+x
				转化为求：”最长可匹配前缀子串“的最大匹配数（已记录在 prefix 表中），然后再比较 最大匹配数 后一个字符是否 == ?
	Kmp 总结
		1.匹配失败时，总是能够让模式串回退到某个位置，使文本不用回退
		2.在字符串比较时，模式串提供的信息越多，计算复杂度越低
	时间复杂度分析：函数 func prefixTableLc(pattern string) []int
		复杂度是线性的（即运算时间与模式串 pattern 的长度是线性关系）
		由 注释2 可以看出 k（最大匹配数） 在整个 for 循环中最多增加 len(pattern) - 1 次，所以让 k 减少的 注释1 在整个 for 循环中最多会执行 len(pattern) - 1 次
		从而 prefixTableLc 函数的复杂度是线性的
	原文参考：灵茶山艾府
		https://www.zhihu.com/question/21923021/answer/37475572
*/

// ====================internet====================

// KmpInternet KMP 字符串匹配算法，网络写法
// prefixTableInternet 所得的前缀表最小值是 -1，故而 j = prefix[j]
func KmpInternet(s, p string) []int {
	n, m := len(s), len(p)-1
	ans := make([]int, 0)
	prefix := prefixTableInternet(p)
	fmt.Println(prefix)
	for i, j := 0, 0; i < n; {
		if j == m && s[i] == p[j] {
			ans = append(ans, i-m)
			j = prefix[j]
		}
		if s[i] == p[j] {
			i++
			j++
		} else {
			j = prefix[j]
			if j == -1 {
				i++
				j++
			}
		}
	}
	return ans
}

func prefixTableInternet(p string) []int {
	//prefix, n := make([]int, len(p)), len(p)-1	// old：n = len(p)-1
	n := len(p)
	prefix := make([]int, n)
	// 如果 i < n-1，则始终 prefix[n-1] = 0，而不是可能的 prefix[n-1] > 0
	// 即匹配到一个 pattern 后，从 j = 0 开始，而不是可能的 j > 0 开始
	for i, length := 2, 0; i < n; {
		if p[i] == p[length] {
			length++
			prefix[i] = length
			i++
		} else {
			if length > 0 {
				length = prefix[length]
			} else {
				i++
			}
		}
	}
	prefix[0] = -1
	return prefix
}

// ====================beauty====================

// KmpBeauty KMP 字符串匹配算法，数据结构与算法之美
// prefixTableBeauty 所得的前缀表最小值是 -1，故而 j = prefix[j-1] + 1
func KmpBeauty(s, p string) []int {
	n, m := len(s), len(p)
	ret, prefix := make([]int, 0), prefixTableBeauty(p)
	fmt.Println(prefix)
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && s[i] != p[j] {
			j = prefix[j-1] + 1
		}
		if s[i] == p[j] {
			j++
		}
		if j == m {
			ret = append(ret, i-m+1)
			j = prefix[j-1] + 1
		}
	}
	return ret
}

func prefixTableBeauty(pattern string) []int {
	k, m := -1, len(pattern)
	prefix := make([]int, m)
	prefix[0] = -1
	for i := 1; i < m; i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			k = prefix[k]
		}
		if pattern[k+1] == pattern[i] {
			k++
		}
		prefix[i] = k
	}
	return prefix
}

// ====================leetcode====================

// KmpLc KMP 字符串匹配算法，LeetCode 灵茶山艾府
// prefixTableLc 所得的前缀表最小值是 0，故而 j = prefix[j-1]
// s：文本字符串，p：模式串，i：文本字符串的索引，j：模式串的索引
// m == 1 时，暴力字符串匹配
func KmpLc(s, p string) []int {
	n, m := len(s), len(p)
	ret, prefix := make([]int, 0), prefixTableLc(p)
	fmt.Println(prefix)
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && p[j] != s[i] {
			j = prefix[j-1]
		}
		if p[j] == s[i] {
			j++
		}
		if j == m {
			ret = append(ret, i-m+1)
			j = prefix[j-1]
		}
	}
	return ret
}

// k：公共前后缀的长度，也指向公共前后缀匹配时前缀的尾字符
// prefix：前缀表。可以实现前缀表的 -1 开头（或最小值）的写法，这里实现 0 开头的写法
func prefixTableLc(pattern string) []int {
	k, m := 0, len(pattern)
	prefix := make([]int, m)
	for i := 1; i < m; i++ {
		for k > 0 && pattern[k] != pattern[i] { // 前缀尾字符和后缀尾字符不相同
			k = prefix[k-1] // 1
		}
		if pattern[k] == pattern[i] { // 前缀尾字符和后缀尾字符相同
			k++ // 2
		}
		prefix[i] = k
	}
	return prefix
}

// ====================扩展 KMP / Z Algorithm====================

// zFunction 对于一个长度为 n 的字符串 s，定义函数 z[i] 表示 s 和 s[i,n-1]（即以 s[i] 开头的后缀）的最长公共前缀（LCP）的长度，则 z 被称为 s 的 Z 函数
// 特别地，z[0] = 0
// 国外一般将计算该数组的算法称为 Z Algorithm，而国内则称其为 扩展 KMP
//
// s[l,r] 是 s 的前缀，即 s[0,r-l] == s[l,r]
// 在计算 z[i] 时我们保证 l <= i，初始时 l=r=0
// 1.核心逻辑：如果 i <= r，那么根据 [l,r] 的定义有 s[i,r] = s[i-l,r-l]，因此 z[i] >= min(z[i-l],r-i+1)
// 1)若 z[i-l] < r-i+1，则 z[i] = z[i-l]
// 2)否则 z[i-l] >= r-i+1，这时我们令 z[i] = r-i+1，然后暴力枚举下一个字符扩展 z[i] 直到不能扩展为止
// 2.如果 i>r，那么我们直接按照朴素算法，从 s[i] 开始比较，暴力求出 z[i]
// 3.在求出 z[i] 后，如果 i+z[i]-1>r，我们就需要更新 [l,r]，即令 l=i, r=i+z[i]-1
//
// 总结：
// s[0,r-l] == s[l,r]，充分利用 s[0,r-l] 的前缀已经计算
// 当 l < i <= r 时，由于 s[i-l,r-l] == s[i,r]，必然有 z[i] >= z[i-l]
// 1.若 z[i-l] < r-i+1，说明 s[i-l,r-l] 不全是前缀匹配，则 z[i] = z[i-l]
// 2.若 z[i-l] >= r-i+1，说明 s[i-l,r-l] 全部前缀匹配，并有未计算的字符匹配，则 z[i] >= r-i+1
func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r && z[i-l] < r-i+1 { // 1
			z[i] = z[i-l]
		} else {
			z[i] = max(0, r-i+1)
			for i+z[i] < n && s[z[i]] == s[i+z[i]] { // 2
				z[i]++
			}
		}
		if i+z[i]-1 > r { // 3
			l, r = i, i+z[i]-1
		}

		// lc 写法：
		z[i] = max(0, min(z[i-l], r-i+1))
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	return z
}
