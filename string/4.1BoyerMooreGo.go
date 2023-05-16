package string

/*
Boyer-Moore 字符串匹配算法
	后缀匹配算法

坏字符规则：
	1.不存在的字符
	2.存在的字符
好后缀规则：
	1.存在与好后缀完全匹配的子串
	2.存在与好后缀部分匹配的子串（前缀）
	3.不存在与好后缀匹配的子串
Go源码 strings.Replacer
	strings.search.go
		makeStringFinder
		next
*/

func makeStringFinder(pattern string) ([256]int, []int) {
	badCharSkip, goodSuffixSkip := [256]int{}, make([]int, len(pattern))
	last := len(pattern) - 1
	for i := range badCharSkip { // 坏字符规则：1.不存在的字符。全部移动 len(pattern) 位
		badCharSkip[i] = len(pattern)
	}
	for i := 0; i < last; i++ { //  // 坏字符规则：2.存在的字符。计算 badCharSkip（从尾第一次出现的位数），不包含最后一位（last := len(pattern) - 1），最后一位移动至少要一位
		badCharSkip[pattern[i]] = last - i
	}
	// Build good suffix table.
	// First pass: set each value to the next index which starts a prefix of
	// pattern.
	// 好后缀表：1
	/*
		补充：2023
			好后缀规则：计算前缀。如 p = "ABABCBAB"，前缀 AB，匹配好后缀 AB（其他的都不匹配）
		第一次循环，比较的是两个空字符串，符合；并且lastPrefix = 字符串的长度
			只有前缀匹配后缀的情况，才会计算新的lastPrefix，例子 p = "ABABCBAB" 中，lastPrefix 只有 "" 和 "AB" 两个前缀符合，则只有 8 和 6 两个值
		lastPrefix：下一次比较的右移距离（last+1 - 前缀的长度）；前缀的长度 = last-i
			则lastPrefix = last+1 - (last-i) = i + 1
		last - i：回移到比较的起始位置（起始位置为p的末尾，则从 i 往后移 last - i）
		前面两者加起来，即为 s 的索引移动的距离，记录到 goodSuffixSkip[]
	*/
	lastPrefix := last           // 回移 i：如 pattern = "ABABCAB"，S = "oooXBABCABooooy"，此时 Pi = 0，Si = 3，goodSuffixSkip[0] = 11(11 = 5 + 6 - 0)，移动到 pattern 尾对齐 y
	for i := last; i >= 0; i-- { // case 2 & 3：模式串存在与好后缀部分匹配的子串 & 模式串不存在与好后缀匹配的子串
		//if HasPrefix(p, p[i+1:]) {
		//len(s) >= len(prefix) && s[0:len(prefix)] == prefix
		if pattern[:last-i] == pattern[i+1:] { // 第一次是两个空字符串
			lastPrefix = i + 1 // lastPrefix：下一次比较的右移距离（last+1 - 前缀的长度）；前缀的长度 = last+1 - (last-i) = i + 1
		}
		// lastPrefix is the sh ift, and (last-i) is len(suffix).
		goodSuffixSkip[i] = lastPrefix + last - i // last-i：回移到比较的起始位置，lastPrefix：下一次比较的右移距离（last+1 - 前缀的长度）
	}
	//fmt.Println("好后缀 1", goodSuffixSkip)

	// Second pass: find repeats of pattern's suffix starting from the front.
	// 好后缀表：2
	/*
		补充：2023
			好后缀规则：计算好后缀。如 p = "ABABCBAB"
		i==last，成了 "ABABCBAB" 和 "BABCBAB" 比较
		lenSuffix = longestCommonSuffix()：计算好后缀的长度
			pattern[1:i+1]：没有截取首字符
		pattern[i-lenSuffix] == pattern[last-lenSuffix]：说明是前缀
			因为longestCommonSuffix已经把相同的字符比较完了，如果还相等，只能是没有截取的首字符
		lenSuffix：作用同上面的lastPrefix
		last - i：同上
		注意：pattern[i-lenSuffix] != pattern[last-lenSuffix]，除了前缀情况，每次都会进入
			显然冗余了，而且冗余的情况都是修改 goodSuffixSkip[0] 的值
			如果 p 结尾不是类似 "...AA" 的情况，goodSuffixSkip[0] 都会是 1
			如 p := "ABABCBABB"，goodSuffixSkip[0] = 2
	*/
	for i := 0; i < last; i++ { // i == last，包含在了前缀表的情况（错）
		lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1]) // case 1：模式串存在与好后缀完全匹配的子串
		//fmt.Println(i, pattern, pattern[1:i+1], lenSuffix)
		//fmt.Println("p:", pattern[1:i+1])
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] { // pattern[i-lenSuffix] == pattern[last-lenSuffix]：说明是前缀
			//fmt.Printf("%c %c\n", pattern[i-lenSuffix], pattern[last-lenSuffix])
			//fmt.Println("last-lenSuffix", last-lenSuffix)
			// (last-i) is the shift, and lenSuffix is len(suffix).
			goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i // last-lenSuffix：修正后缀表的索引，lenSuffix：回移到比较的起始位置，last - i：下一次比较的右移距离
		}
	}
	//fmt.Println("好后缀 2", goodSuffixSkip)
	return badCharSkip, goodSuffixSkip
}

// 仅比较两个字符串的共同后缀长度，没有则为0
func longestCommonSuffix(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] { // 从后往前比较，寻找与好后缀完全匹配的子串
			break
		}
	}
	return
}
func next(pattern, text string) []int {
	badCharSkip, goodSuffixSkip := makeStringFinder(pattern)
	//fmt.Println(badCharSkip, goodSuffixSkip)
	ans, n := make([]int, 0), len(pattern)-1
	for i := n; i < len(text); { // text的索引
		j := n                                // pattern的索引
		for j >= 0 && text[i] == pattern[j] { // 从后往前，开始匹配
			i--
			j--
		}
		if j < 0 { // 匹配成功
			//fmt.Println("i,j:", i, j)
			i, j = i+1, 0 // 重置 i j，匹配下一个
			ans = append(ans, i)
		}
		//fmt.Println(i, badCharSkip[text[i]], j, goodSuffixSkip[j])
		i += max(badCharSkip[text[i]], goodSuffixSkip[j]) // 根据好后缀、坏字符规则，挪动 i
		//fmt.Println(i, j)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
