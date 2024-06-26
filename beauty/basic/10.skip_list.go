package basic

/*
跳表：为什么Redis一定要用跳表来实现有序集合？

跳表（Skip list）
	索引/索引层
	链表加多级索引的结构，就是跳表
	使用空间换时间的设计思路，通过构建多级索引来提高查询的效率，实现了基于链表的“二分查找”

	在链表中使用二分查找
	动态数据结构，甚至可以替代红黑树（Red-black tree）
	Redis 中的有序集合（Sorted Set）就是用跳表来实现的

用跳表查询到底有多快？
	第 k 级索引的结点个数是第 k-1 级索引的结点个数的 1/2，那第 k 级索引结点的个数就是 n/(2^k)
	假设索引有 h 级，最高级的索引有 2 个结点。则 n/(2^h)=2，从而求得 h=log2(n)-1
	包含原始链表这一层，整个跳表的高度就是 log2(n)
	在跳表中查询某个数据的时候，如果每一层都要遍历 m 个结点，那在跳表中查询一个数据的时间复杂度就是 O(m*log n)
		每两个结点会抽出一个结点作为上一级索引的结点，m 值为 3：每一级最多遍历3个结点
	O(log n)
跳表是不是很浪费内存？
	空间换时间
	索引需要 O(n) 额外空间：n/2+n/4+n/8…+8+4+2=n-2
	实际上，在软件开发中，不必太在意索引占用的额外空间
	在实际的软件开发中，原始链表中存储的有可能是很大的对象（而非整数），而索引结点只需要存储关键值和几个指针，并不需要存储对象
高效的动态插入和删除
	查找，再插入：O(log n)
	查找前驱节点，再删除：O(log n)

跳表索引动态更新
	通过随机函数来维护前面提到的“平衡性”
	随机函数决定将这个结点插入到哪几级索引中，比如随机函数生成了值 K，那我们就将这个结点添加到第一级到第 K 级这 K 级索引中
	随机函数的选择很有讲究，从概率上来讲，能够保证跳表的索引大小和数据大小平衡性，不至于性能过度退化
跳表的实现
	https://github.com/wangzheng0822/algo

Redis 为什么会选择用跳表来实现有序集合呢？ 为什么不用红黑树呢？
	Redis 中的有序集合是通过跳表来实现的，严格点讲，其实还用到了散列表
	Redis 中的有序集合支持的核心操作
		插入一个数据
		删除一个数据
		查找一个数据
		按照区间查找数据（比如查找值在[100, 356]之间的数据）
		迭代输出有序序列
	红黑树
		插入、删除、查找以及迭代输出有序序列这几个操作，红黑树也可以完成，时间复杂度跟跳表是一样的
		跳表也不能完全替代红黑树。因为红黑树比跳表的出现要早一些，很多编程语言中的 Map 类型都是通过红黑树来实现的
		业务开发的时候，直接拿来用
	跳表
		跳表可以做到 O(logn) 的时间复杂度定位区间的起点，然后在原始链表中顺序往后遍历
		跳表更容易代码实现
		跳表更加灵活，可以通过改变索引构建策略，有效平衡执行效率和内存消耗

小结
	跳表使用空间换时间的设计思路，通过构建多级索引来提高查询的效率，实现了基于链表的“二分查找”
	跳表是一种动态数据结构，支持快速地插入、删除、查找操作，时间复杂度都是 O(logn)

思考
	如果每三个或者五个结点提取一个结点作为上级索引，对应的在跳表中查询数据的时间复杂度是多少呢？
		x 个结点：O((x+1) log (x,n))
*/
