package basic

/*
堆和堆排序：为什么说堆排序没有快速排序快？

“堆”（Heap）
	堆是一个完全二叉树
	堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值
	大顶堆
	小顶堆
如何实现一个堆？
	往堆中插入一个元素
		堆化（heapify）：从下往上 / 从上往下
	删除堆顶元素
如何基于堆实现排序？
	建堆：两种思路
		在堆中插入一个元素的思路，从下往上堆化
		从后往前处理数组，从上往下堆化，叶子节点不需要堆化
	建堆时间复杂度：O(n)
		每个节点堆化的过程中，需要比较和交换的节点个数，跟这个节点的高度 k 成正比
		S1 = 1*h + 2*(h-1) + 2^2*(h-2) + ... + 2^(h-1) * 1
		S2 = 2*S1
		S = S2-S1 = -h + 2 + 2^2 + ... + 2^(h-1) + 2^h
		S = 2^(h+1)-h-2
		h = log2(n)，带入公式 S，则 S = O(n)
		详见：https://time.geekbang.org/column/article/69913
	排序
		原地排序算法
		O(n log n)
		不稳定排序算法

在实际的软件开发中，快速排序的性能要比堆排序好，这是为什么呢？
	堆排序数据访问的方式没有快速排序友好
		对于快速排序来说，数据是顺序访问的
		而对于堆排序来说，数据是跳着访问的，对 CPU 缓存是不友好的
	对于同样的数据，在排序过程中，堆排序算法的数据交换次数要多于快速排序
		快速排序数据交换的次数不会比逆序度多
		堆排序比快速排序交换次数多
			建堆的过程会打乱数据原有的相对先后顺序，导致原数据的有序度降低
			对于一组已经有序的数据来说，经过建堆之后，数据反而变得更无序了

思考
	在堆排序建堆的时候，对于完全二叉树来说，下标从 n/2+1 到 n 的都是叶子节点，这个结论是怎么推导出来的呢？
		1.反证法：i>>1 如果有数据，那 i 不是叶子节点
		2.补全为满二叉树，少了 x 个最深层的结点，则多了 x>>2 个上一层叶子结点
		3.递推公式：叶子节点总数 x+y，则非叶子节点的总数为 (x+1)>>1 + 2^k + 2^(k-1) + ... + 1，且 (x+1)>>1 + y>>1 == 2^k
			...
	堆的一种经典应用，堆排序。关于堆，还能想到它的其他应用吗？
		topK
		中位数（对顶堆）
		优先队列
*/
