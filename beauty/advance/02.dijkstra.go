package advance

/*
最短路径：地图软件是如何计算出最优出行路径的？

针对无权图的搜索算法
	深度优先搜索和广度优先搜索
地图软件的最优路线是如何计算出来的吗？底层依赖了什么算法呢？
	最短路线：最短路径算法（Shortest Path Algorithm），单源最短路径算法
	最少用时路线
	最少红绿灯路线
	...

算法解析
	建模：解决软件开发中的实际问题，最重要的一点就是建模，将复杂的场景抽象成具体的数据结构
		把地图抽象成图（有向有权图）
		每个岔路口看作一个顶点，岔路口与岔路口之间的路看作一条边，路的长度就是边的权重
		如果路是单行道，就在两个顶点之间画一条有向边
		如果路是双行道，我们就在两个顶点之间画两条方向不同的边

		问题就转化为，在一个有向有权图中，求两个顶点间的最短路径
	最短路径算法：单源最短路径算法（一个顶点到一个顶点）
		Dijkstra算法：最有名的最短路径算法
		Bellford 算法
		Floyd 算法
Dijkstra
	实现思路
		用一个优先级队列，来记录已经遍历到的顶点以及这个顶点与起点的路径长度
	adj：邻接表
	queue：优先队列，小顶堆
		记录当前能到达的所有顶点，优先级为这些顶点与起始顶点的当前最短路径
		最核心的算法
	minVertex：优先级队列中 dist 最小的顶点
	nextVertex：minVertex 可达的所有顶点
		minVertex.dist + w < nextVertex.dist 时
		更新 nextVertex.dist = minVertex.dist + w
		w 为 minVertex 与 nextVertex 之间边的权重 w
	vertexes 数组：记录从起始顶点到每个顶点的距离（dist）
	predecessor 数组：作用是为了还原最短路径，它记录每个顶点的前驱顶点
	inqueue 数组：避免将一个顶点多次添加到优先级队列中

	示例：参见 02.dijkstra.jpg
	时间复杂度：O(E*logV)
		for 循环最多执行次数为边的个数 E
		优先队列的操作时间复杂度是 O(log V)，V 是顶点的个数
	代码实现
		E:\gothmslee\algo\algo\dijkstra.go

Dijkstra 工程实践
	理论缺点
		对于一个超级大地图来说，岔路口、道路都非常多，对应到图这种数据结构上来说，就有非常多的顶点和边
		如果为了计算两点之间的最短路径，在一个超级大图上动用 Dijkstra 算法，遍历所有的顶点和边，显然会非常耗时
	工程实践原则
		做工程不像做理论，一定要给出个最优解，理论上算法再好，如果执行效率太低，也无法应用到实际的工程中
		要根据问题的实际背景，对解决方案权衡取舍。类似出行路线这种工程上的问题，没有必要非得求出个绝对最优解
		很多时候，为了兼顾执行效率，我们只需要计算出一个可行的次优解就可以了
	方案思路
		两点较近
			两点之间的最短路径或者说较好的出行路径，并不会很“发散”，只会出现在两点之间和两点附近的区块内
			可以在整个大地图上，划出一个小的区块，这个小区块恰好可以覆盖住两个点，但又不会很大
			只需要在这个小区块内部运行 Dijkstra 算法，这样就可以避免遍历整个大图，也就大大提高了执行效率
		两点较远：如从北京海淀区某个地点，到上海黄浦区某个地点
			可以把北京海淀区或者北京看作一个顶点，把上海黄浦区或者上海看作一个顶点
			先规划大的出行路线
				比如，如何从北京到上海，必须要经过某几个顶点，或者某几条干道
			后再细化每个阶段的小路线
	最少时间
		把边的权重，从路的长度变成经过这段路所需要的时间
		但是时间会根据拥堵情况时刻变化
		如何计算车通过一段路的时间呢？
	最少红绿灯
		假设每经过一条边，就要经过一个红绿灯
		则只需要把每条边的权值改为 1 即可，算法还是不变
		BFS：
			但是，边的权值为 1，也就相当于无权图了，我们还可以使用之前讲过的广度优先搜索算法
			广度优先搜索算法计算出来的两点之间的路径，就是两点的最短路径
	地图软件的路径规划
		真实的地图软件的路径规划，要比这个复杂很多
		比起 Dijkstra 算法，地图软件用的更多的是类似 A* 的启发式搜索算法，不过也是在 Dijkstra 算法上的优化罢了

翻译系统
	场景描述
		系统：
			有一个翻译系统，只能针对单个词来做翻译
			如果要翻译一整个句子，我们需要将句子拆成一个一个的单词，再丢给翻译系统
			针对每个单词，翻译系统会返回一组可选的翻译列表，并且针对每个翻译打一个分，表示这个翻译的可信程度
		翻译：
			针对每个单词，从可选列表中，选择其中一个翻译，组合起来就是整个句子的翻译
			每个单词的翻译的得分之和，就是整个句子的翻译得分
			随意搭配单词的翻译，会得到一个句子的不同翻译
			针对整个句子，我们希望计算出得分最高的前 k 个翻译结果，你会怎么编程来实现呢？
	回溯算法
		穷举所有的排列组合情况，然后选出得分最高的前 k 个翻译结果
		时间复杂度：O(m^n)，m 表示平均每个单词的可选翻译个数，n 表示一个句子中包含多少个单词
	借助 Dijkstra 算法的核心思想
		a0b0c0 肯定是得分最高组合结果，放入到优先级队列中
		每次从优先级队列中取出一个得分最高的组合，并基于这个组合进行扩展
			扩展的策略是每个单词的翻译分别替换成下一个单词的翻译
			a0b0c0 扩展得到三个组合，a1b0c0 a0b1c0 a0b0c1，加到优先级队列中
		重复这个过程，直到获取到 k 个翻译组合或者队列为空

		图解：参见 02.dijkstra-translate-01.jpg & 02.dijkstra-translate-02.jpg
		时间复杂度：
			O(k*n*log(k*n))
				包含 n 个单词，求得分最高的前 k 个组合结果，X 表示队列中的组合个数，的时间复杂度就是 O(k*n*logX)
			X：k 次出队列，队列中的总数据不会超过 k*n（每次有一个组合出队列，就有 n 个组合入队列）
				出队、入队操作的时间复杂度是 O(log(k*n))

思考
	1.在计算最短时间的出行路线中，如何获得通过某条路的时间呢？这个题目很有意思，我之前面试的时候也被问到过，你可以思考看看
		// TODO
	2.出行路线问题，假设的是开车出行，那如果是公交出行呢？如果混合地铁、公交、步行，又该如何规划路线呢？

	参考
	1.获取通过某条路的时间：通过某条路的时间与①路长度②路况(是否平坦等)③拥堵情况④红绿灯个数等因素相关。获取这些因素后就可以建立一个回归模型(比如线性回归)来估算时间。其中①②④因素比较固定，容易获得。③是动态的，但也可以通过a.与交通部门合作获得路段拥堵情况；b.联合其他导航软件获得在该路段的在线人数；c.通过现在时间段正好在次路段的其他用户的真实情况等方式估算。

	2.混合公交、地铁和步行时：地铁时刻表是固定的，容易估算。公交虽然没那么准时，大致时间是可以估计的，步行时间受路拥堵状况小，基本与道路长度成正比，也容易估算。总之，感觉公交、地铁、步行，时间估算会比开车更容易，也更准确些。
*/
