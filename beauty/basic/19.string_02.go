package basic

/*
字符串匹配基础（中）：如何实现文本编辑器中的查找功能？

BM（Boyer-Moore）算法
	是著名的 KMP 算法的 3 到 4 倍
	核心思想：
		利用模式串本身的特点，在模式串中某个字符与主串不能匹配的时候，将模式串往后多滑动几位，以此来减少不必要的字符比较，提高匹配的效率
BM 算法原理分析
	坏字符规则（bad character rule）
	好后缀规则（good suffix shift）
坏字符规则
	从模式串的末尾往前倒着匹配，当发现某个字符没法匹配的时候，这个没有匹配的字符叫作坏字符（主串中的字符）
	如果坏字符在模式串里多处出现，那选择最靠后的那个
	最好情况下的时间复杂度非常低，是 O(n/m)
好后缀规则
	这里仅截取规则示意部分，具体计算参考 string.4.2BoyerMoore.go
	规则一：对应 好后缀规则2，计算前缀
		好后缀 "AB" 存在与之匹配的前缀 "AB"，则 prefixIdx = 6
			  最后一个AB            第一个AB             好后缀AB的起始索引
		第二个 "AB" 不是前缀，前缀必须是从 0 位置开始（但是它符合下面2.3的情况）
			示例：p = "ABABCBAB"
			依次比较：前缀 vs p
			"A" vs "B"：lastIdx - i + prefixIdx = 7 - 6 + 8
			"AB" vs "AB"：lastIdx - i + prefixIdx = 7 - 5 + 6
			...
			"ABABCBA" vs "BABCBAB"：lastIdx - i + prefixIdx = 7 - 0 + 6
	规则二：对应 好后缀规则1，计算好后缀
		last：len(p) - 1
		j：计算完全匹配字符串的长度（即好后缀的长度，suffix len）
		j：回移到匹配的起始位置
		lastIdx - i：下一次匹配的右移距离
		lastIdx-j：修正后缀表的索引
		计算与好后缀匹配的子串（非前缀）:
			示例：p = "ABABCBAB"，j保存完全匹配字符串的长度（倒序遍历）
			"" vs ABABCBAB：j = 0
			B vs ABABCBAB：j = 1, i = 1, lastIdx - i = 6, lastIdx-j = 6：前缀的情况
			BA vs ABABCBAB：j = 0
			BAB vs ABABCBAB：j = 3, i = 3, lastIdx - i = 4, lastIdx-j = 4：goodSuffix[4] 修正为 7（goodSuffix[lastIdx-j] = j + lastIdx - i）
			BABC vs ABABCBAB：j = 0
			BABCB vs ABABCBAB：j = 1, i = 5, lastIdx - i = 2, lastIdx-j = 6：如 BABCB 最右边的 B 移动到好后缀 B，需要移动 2 个位置，修正 goodSuffix[6] = 1 + 2
			BABCBA vs ABABCBAB：j = 0

BM 算法代码实现
	实现 badCharSkip
	实现 goodSuffixSkip：核心内容	// TODO
		在模式串中，查找跟好后缀匹配的另一个子串
			好后缀规则
		在好后缀的后缀子串中，查找最长的、能跟模式串前缀子串匹配的后缀子串
			前缀规则
	实现 next 匹配函数
BM 算法的性能分析及优化
	舍弃坏字符规则：
		如果处理字符集很大的字符串匹配问题，bc 数组对内存的消耗就会比较多
		因为好后缀和坏字符规则是独立的，如果我们运行的环境对内存要求苛刻，可以只使用好后缀规则，不使用坏字符规则，这样就可以避免 bc 数组过多的内存消耗
		不过，单纯使用好后缀规则的 BM 算法效率就会下降一些
	初级版本：
		专栏中 BM 的实现是个初级版本，尤其是 goodSuffixSkip 的实现
		“A new proof of the linearity of the Boyer-Moore string searching algorithm”证明了在最坏情况下，BM 算法的比较次数上限是 5n
			https://dl.acm.org/doi/10.1109/SFCS.1977.3
			证明了在最坏情况下，BM 算法的比较次数上限是 5n
		“Tight bounds on the complexity of the Boyer-Moore string matching algorithm”证明了在最坏情况下，BM 算法的比较次数上限是 3n
			https://dl.acm.org/doi/10.5555/127787.127830
			证明了在最坏情况下，BM 算法的比较次数上限是 3n

小结
	BM 算法核心思想
		利用模式串本身的特点，在模式串中某个字符与主串不能匹配的时候，将模式串往后多滑动几位，以此来减少不必要的字符比较，提高匹配的效率
	好后缀规则可以独立于坏字符规则使用
		因为坏字符规则的实现比较耗内存，为了节省内存，我们可以只用好后缀规则来实现 BM 算法

思考
	对于查找功能是重要功能的软件来说，比如一些文本编辑器，它们的查找功能都是用哪种算法来实现的呢？
		补充：有没有比 BF 算法和 RK 算法更加高效的字符串匹配算法呢？
	熟悉的编程语言中的查找函数，或者工具、软件中的查找功能，都是用了哪种字符串匹配算法呢？
		暴力字符串匹配
		Boyer-Moore
		Rabin-Karp
*/
