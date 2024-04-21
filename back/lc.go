package main

/*
vs DP
	第 139 题可以使用动态规划的方法判断是否可以拆分，因此这道题也可以使用动态规划的思想。但是这道题如果使用自底向上的动态规划的方法进行拆分，则无法事先判断拆分的可行性，在不能拆分的情况下会超时
	例如以下例子，由于字符串 s 中包含字母 b，而单词列表 wordDict 中的所有单词都由字母 a 组成，不包含字母 b，因此不能拆分，但是自底向上的动态规划仍然会在每个下标都进行大量的匹配，导致超时
		s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		wordDict = ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
	为了避免动态规划的方法超时，需要首先使用第 139 题的代码进行判断，在可以拆分的情况下再使用动态规划的方法进行拆分。相比之下，自顶向下的记忆化搜索可以在搜索过程中将不可以拆分的情况进行剪枝，因此记忆化搜索是更优的做法

回溯问题的时间复杂度，通用公式：
	路径长度×搜索树的叶子数
	如 lc-216：O(k C(9,k))，时间/长度为 k，组合数为 C(9,k)

lc
	39：选/不选，枚举，完全背包预处理+可行性剪枝
		https://leetcode.cn/problems/combination-sum/solutions/2747858/liang-chong-fang-fa-xuan-huo-bu-xuan-mei-mhf9/
	216
	140
	93
*/
