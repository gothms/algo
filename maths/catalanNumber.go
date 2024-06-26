package maths

/*
卡特兰数：
	1.计算某个长度为 n 的入栈序列可以有多少种出栈序列
	2.计算包含 n 个节点的二叉树有多少种形状
		lc-96
	两道题的结果是相等的
		每种二叉树形状的中序遍历都能够对应上一种出栈顺序
卡特兰数公式：C(1) = 1，C(n) = C(n-1) * (n<<2-2) / (n+1)
	递推公式：f(n) = C(2n,n) / (n+1)
	递推关系：f(n) = C(2n,n) - C(2n,n-1)
	百科：https://baike.baidu.com/item/%E5%8D%A1%E7%89%B9%E5%85%B0%E6%95%B0/6125746?fromtitle=catalan&fromid=7605685
	以进出栈为例
	常规分析：f(0),f(1)=1,1
	非常规分析：把进栈设为状态 1，出栈设为状态 0，n个数的所有状态对应n个1和n个0组成的2n位二进制数
			输出序列的总数目 = 由左而右扫描由n个1和n个0组成的2n位二进制数，1的累计数不小于0的累计数的方案种数
		C(2n,n)：在2n位二进制数中填入n个1的方案数，不填1的其余n位自动填0
		C(2n,n-1)：不符合要求的方案数
			不符合要求的数的特征，必在某一奇数位2m+1位上首先出现m+1个0的累计数和m个1的累计数，此后的2(n-m)-1位上有n-m个1和n-m-1个0
			若把后面这2(n-m)-1位上的0和1互换，使之成为n-m个0和n-m-1个1，结果得1个由n+1个0和n-1个1组成的2n位数，即一个不合要求的数对应于一个由n+1个0和n-1个1组成的排列
			反过来，任何一个由n+1个0和n-1个1组成的2n位二进制数，由于0的个数多2个，2n为偶数，故必在某一个奇数位上出现0的累计数超过1的累计数
			同样在后面部分0和1互换，使之成为由n个0和n个1组成的2n位数，即n+1个0和n-1个1组成的2n位数必对应一个不符合要求的数
		输出序列的总数目 = C(2n,n)-C(2n,n-1)=C(2n,n)/(n+1) = f(n)
参考
	E:\gothmslee\algo\beauty\basic\07.recursive.go
*/
