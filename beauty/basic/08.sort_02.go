package basic

/*
归并排序原理（Merge Sort）
	分治思想
哨兵
	利用哨兵，简化 merge 函数

分治
	分而治之，将一个大问题分解成小的子问题来解决
分治 vs 递归
	分治算法一般都是用递归来实现的
	分治是一种解决问题的处理思想，递归是一种编程技巧

归并排序的性能分析
	归并排序是稳定的排序算法吗？
		稳定排序算法
	归并排序的时间复杂度是多少？
		不仅递归求解的问题可以写成递推公式，递归代码的时间复杂度也可以写成递推公式
		利用递归公式，求出归并排序的时间复杂度
			T(1) = C：n=1时，只需要常量级的执行时间，所以表示为C
			T(n) = 2*T(n/2) + n：n>1
	归并排序的空间复杂度是多少？
		1.不是原地排序算法
		2.通过递推公式来求解，那整个归并过程需要的空间复杂度就是 O(nlogn)
		3.在任意时刻，CPU 只会有一个函数在执行，也就只会有一个临时的内存空间在使用
			临时内存空间最大也不会超过 n 个数据的大小，所以空间复杂度是 O(n)

快速排序原理（Quicksort）
	分治思想

快速排序的性能分析
	快速排序是稳定的排序算法吗？
		非稳定排序算法
	快速排序的时间复杂度是多少？
		最坏：O(n^2)，数组原来已经有序
		递归树求解递归的时间复杂度
		// TODO
	快速排序的空间复杂度是多少？
		原地排序算法

归并排序 vs 快速排序
	处理过程
		归并：由下到上，先处理子问题，再合并
		快排：由上到下，先分区，再处理子问题
	是否原地
		归并：非原地排序算法
		快排：原地排序算法

如何在 O(n) 的时间复杂度内查找一个无序数组中的第 K 大元素？
	每次分区遍历元素个数的和：n+n/2+n/4+...+1 = 2n-1
	则快排解决 TopK 问题时间复杂度是 O(n)

思考
	现在你有 10 个接口访问日志文件，每个日志文件大小约 300MB，每个文件里的日志都是按照时间戳从小到大排序的
	你希望将这 10 个较小的日志文件，合并为 1 个日志文件，合并之后的日志仍然按照时间戳从小到大排列
	如果处理上述排序任务的机器内存只有 1GB，你有什么好的解决思路，能“快速”地将这 10 个日志文件合并吗？
*/