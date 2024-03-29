package beauty

/*
画图工具
	iPad Paper

1.对于 advance.06.b+.go 思考 2. 问题
	对平衡二叉查找树进行改造，将叶子节点串在链表中，就支持了按照区间来查找数据
	散列表也经常跟链表一块使用，如果我们把散列表中的结点，也用链表串起来，能否支持按照区间查找数据呢？

	问题：散列冲突
		假如字符串 str1 str2 ... strn 的 sha1 值都为 xxx
		那么 xxx+str1/str1+str1 xxx+str2 ... xxx+strn 的 sha1 值相同的概率有多大呢？
		有没有一种从“根本”上解决散列冲突的方法？
*/
