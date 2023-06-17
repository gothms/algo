package string

/*
KMP 字符串匹配算法
	前缀匹配算法

前缀表：prefix-table
	最长公共前后缀
	示例：p = "ABABCABAA" prefix = [0 0 1 2 0 1 2 3 1]
		sub 为模式串 p [0,i] 的子串，i<=len(p)-1
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
	前缀表优化：后移一位
		为什么要后移一位？
		1.匹配 s[i] =? p[j] 时，若不匹配，查询的是 prefix[j-1]，所以 prefix 最后一个元素其实是无用处的
		2.模式串p 和 前缀表prefix 的索引 j 和 j-1 不对应，后移一位正好对应
			同时将 prefix[0] 置为 -1，让计算更简便
代码注释：
	1.前缀表：可以实现前缀表的 0 开头的写法，这里实现 -1 开头的写法
		p：模式串
		s：文本字符串
		i：文本字符串的索引
		j：公共前后缀的长度，j 也指向公共前后缀的下一个字符
		1.1.前缀尾字符和后缀尾字符相同
			prefix[i] = j+1：即 i-1 时，最长公共前后缀 + 1
			示例：p = "ABABCABAA" prefix = [0 0 1 2 0 1 2 3 1]
				i =          7
				j =     2
				p[i] == p[j]
				则有：prefix[7] = j+1 = 3
		1.2.两者尾字符不同
			1.2.1.先重置 j（公共前后缀的长度）的值，再计算（难点）
				示例：p := "abaababaab" dp = [0 0 1 1 2 3 2 3 4 5]
				j-1：前缀的最后一个字符
				dp[j-1]：前缀的最后一个字符位置，有多长的公共前后缀
				示例：p := "abaababaab"
					i =          6：当前计算的子串为 abaabab
					j =       3：i=5时，公共前后缀的长度为 3（j 也指向公共前后缀的下一个字符）
					j-1 =    2：i=5时，公共前缀为 aba
					dp[j-1]=1：表示
						p := "abaababaab"
							  aba
						aba 公共前后缀的长度为 1
					j = prefix[j-1]：j = 1
						p := "abaababaab"
							  ab
						b 即为现在需要比较的字符：p[i] =? p[1]
			1.2.2.上一个最长公共前后缀长度为 0，则当前计算的最长公共前后缀长度也为 0：省略 dp[i] = 0
				i++，进行下一次计算

			补充：j两个作用，1.公共前后缀的长度，2.公共前后缀结尾字符的索引（或索引的左边字符），这点非常重要
				1.2.1.理解length = prefix[length-1]
					例，ABADABAAC，D!=A时，i=7，length=3，prefix[7]的取值区间为[0,3]
					如果ABA=BAA，为3；如果AB=AA，为2；如果A=A，为1。前两种情况，都可以推出prefix[3-1]>0
					因为prefix[6]=3，则找ABAA可能存在的公共前后缀，即为找首3个字符，再加上p[7]的公共前后缀
					所以prefix[3-1]>0，则prefix[7]可能为[0,3]

					上述描述拗口，不如画图理解，后续补上
				1.2.2.假如length=0，则比较的是p[0]==p[i]
					前面已经比较不相同了，所以prefix[i]=0；则比较下一个，i++
	2.KMP搜索：
		2.1.匹配成功
			2.1.1.len(p) == 1 时，j 为 -1
			2.1.2.重置 i j 进行，开始新的一次匹配
				也可以打补丁 if len(p) == 1：使用蛮力搜索匹配字符串
			2.1.3.避免 len(p] == 1，出现 j 越界
		2.2.匹配失败
			2.2.1.根据前缀表，移动模式串
			2.2.2.匹配失败的是模式串的第一个字符
				重置 i j 进行，开始新的一次匹配
*/

// Kmp 字符串匹配算法
func Kmp(s, p string) []int {
	n, m := len(s), len(p)-1
	ans := make([]int, 0)
	prefix := prefixTable(p)   // 1
	for i, j := 0, 0; i < n; { // 3
		if j == m && s[i] == p[j] { // 3.1
			ans = append(ans, i-m)
			j = prefix[j]
		}
		if s[i] == p[j] {
			i++
			j++
		} else { // 3.2
			j = prefix[j]
			if j == -1 {
				i++
				j++
			}
		}
	}
	return ans
}

func prefixTable(p string) []int {
	prefix, n := make([]int, len(p)), len(p)-1
	for i, length := 2, 0; i < n; i++ {
		if p[i] == p[length] { // 1.1
			length++
			prefix[i] = length
			i++
		} else { // 1.2
			if length > 0 { // 1.2.1
				length = prefix[length]
			} else { // 1.2.2
				i++
			}
		}
	}
	prefix[0] = -1
	return prefix
}
