package basic

/*
贪心算法：如何用贪心算法实现Huffman压缩编码？

贪心算法（greedy algorithm）
	霍夫曼编码（Huffman Coding）
	Prim
	Kruskal 最小生成树算法
	Dijkstra 单源最短路径算法
霍夫曼编码
	利用贪心算法来实现对数据压缩编码，有效节省数据存储空间

如何理解“贪心算法”？
	第一步，当我们看到这类问题的时候，首先要联想到贪心算法
		针对一组数据，我们定义了限制值和期望值，希望从中选出几个数据，在满足限制值的情况下，期望值最大
	第二步，我们尝试看下这个问题是否可以用贪心算法解决
		每次选择当前情况下，在对限制值同等贡献量的情况下，对期望值贡献最大的数据
	第三步，我们举几个例子看下贪心算法产生的结果是否是最优的
		严格地证明贪心算法的正确性，是非常复杂的，需要涉及比较多的数学推理

	实际上，用贪心算法解决问题的思路，并不总能给出最优解
		贪心算法不工作的主要原因是，前面的选择，会影响后面的选择
		贪心算法的最难的一块是如何将要解决的问题抽象成贪心算法模型，只要这一步搞定之后，贪心算法的编码一般都很简单

贪心算法实战分析
	分糖果
		每次从剩下的孩子中，找出对糖果大小需求最小的，然后发给他剩下的糖果中能满足他的最小的糖果
	钱币找零
		先用面值最大的来支付，如果不够，就继续用更小一点面值的，以此类推，最后剩下的用 1 元来补齐
	区间覆盖：从 n 个区间中选出一部分区间，这部分区间满足两两不相交（端点相交的情况不算相交），最多能选出多少个区间呢？
		假设这 n 个区间中最左端点是 lmin，最右端点是 rmax
		每次选择的时候，左端点跟前面的已经覆盖的区间不重合的，右端点又尽量小的
霍夫曼编码
	假设有一个包含 1000 个字符的文件，每个字符占 1 个 byte（1byte=8bits），存储这 1000 个字符就一共需要 8000bits
	那有没有更加节省空间的存储方式呢？
	常规方法：
		通过统计分析发现，这 1000 个字符中只包含 6 种不同字符，假设它们分别是 a、b、c、d、e、f。而 3 个二进制位（bit）就可以表示 8 个不同的字符
		所以，为了尽量减少存储空间，每个字符我们用 3 个二进制位来表示。那存储这 1000 个字符只需要 3000bits 就可以了
		编码：
			a(000)、b(001)、c(010)、d(011)、e(100)、f(101)
		解压：
			对于等长的编码来说，次从文本中读取 3 位二进制码，然后翻译成对应的字符
	霍夫曼编码：
		霍夫曼编码是一种十分有效的编码方法，广泛用于数据压缩中，其压缩率通常在 20%～90% 之间
			1000 个字符只需要 2100bits 就可以了
		霍夫曼编码不仅会考察文本中有多少个不同字符，还会考察每个字符出现的频率，根据频率的不同，选择不同长度的编码
			出现频率比较多的字符，用稍微短一些的编码
			出现频率比较少的字符，用稍微长一些的编码
		编码：
			a、b、c、d、e、f 编码 1、01、001、0001、00001、00000
		解压：
			对于不等长的编码来说，为了避免解压缩过程中的歧义，霍夫曼编码要求各个字符的编码之间，不会出现某个编码是另一个编码前缀的情况
			每次会读取尽可能长的可解压的二进制串，所以在解压缩的时候也不会歧义
如何根据字符出现频率的不同，给不同的字符进行不同长度的编码呢？
	a、b、c、d、e、f 出现频率 450、350、90、60、30、20a、b、c、d、e、f
		详见 Huffman-Coding.jpg
	1.把每个字符看作一个节点，并且附带着把频率放到优先级队列中
		从队列中取出频率最小的两个节点 A、B，然后新建一个节点 C，把频率设置为两个节点的频率之和，并把这个新节点 C 作为节点 A、B 的父节点
		最后再把 C 节点放入到优先级队列中
		重复这个过程，直到队列中没有数据
	2.给每一条边加上画一个权值，指向左子节点的边我们统统标记为 0，指向右子节点的边，我们统统标记为 1
		那从根节点到叶节点的路径就是叶节点对应字符的霍夫曼编码
	补充
		霍夫曼编码构造的树型结构，也用于解码
压缩算法扩展
	https://mp.weixin.qq.com/s/VDjLpgGP233-T0Df-w3s2A

小结
	贪心算法的最难的一块是如何将要解决的问题抽象成贪心算法模型，只要这一步搞定之后，贪心算法的编码一般都很简单
	很多时候，只需要多举几个例子，看一下贪心算法的解决方案是否真的能得到最优解就可以了

思考
	在一个非负整数 a 中，希望从中移除 k 个数字，让剩下的数字值最小，如何选择移除哪 k 个数字呢？
		从高位到低位，将数字维护为：非单调递减
		i：标记当前位，更高位为 i+1
		i 位数字 > i-1 位数字：移除 i，若存在 i+1，i=i+1。否则 i=i-1
		i 位数字 <= i-1 数字：i=i-1
		最后没移够 k，则从最低位开始往高位一个个移除
	假设有 n 个人等待被服务，但是服务窗口只有一个，每个人需要被服务的时间长度是不同的，如何安排被服务的先后顺序，才能让这 n 个人总的等待时间最短？
		将队列排序为：非单调递减
*/
