package basic

/*
字符串匹配基础（下）：如何借助BM算法轻松理解KMP算法？

KMP 算法基本原理
	根据三位作者（D.E.Knuth，J.H.Morris 和 V.R.Pratt）的名字来命名的，算法的全称是 Knuth Morris Pratt 算法
	核心思想：
		在模式串与主串匹配的过程中，当遇到不可匹配的字符的时候，我们希望找到一些规律，可以将模式串往后多滑动几位，跳过那些肯定不会匹配的情况
	不能匹配的那个字符仍然叫作坏字符
	已经匹配的那段字符串叫作好前缀
	next 数组
		失效函数（failure function）
前缀规则
	在模式串和主串匹配的过程中，当遇到坏字符后，对于已经比对过的好前缀，能否找到一种规律，将模式串一次性滑动很多位？
	最长可匹配后缀子串
		好前缀的所有后缀子串中，最长的可匹配前缀子串的那个后缀子串
	最长可匹配前缀子串
		对应的前缀子串
前缀表示例
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

失效函数（failure function）计算方法
	原文中描述比较拗口，参考：https://www.zhihu.com/question/21923021/answer/37475572
	实现：E:\gothmslee\algo\string\3.KMP.go

KMP 算法复杂度分析
	空间复杂度：O(m)，m 表示模式串的长度
	时间复杂度：O(m+n)

思考
	BF 算法、RK 算法、BM 算法和 KMP 算法，关于这些算法，什么地方最难理解呢？
		KMP：k = next[k]
			因为前一个的最长串的下一个字符不与最后一个相等，需要找前一个的次长串，问题就变成了求 0 到 next(k) 的最长串
			如果下个字符与最后一个不等，继续求次长串，也就是下一个next(k)，直到找到，或者完全没有
		参考：E:\gothmslee\algo\string\3.KMP.go
*/
