package basic

/*
排序（下）：如何用快排思想在O(n)内查找第K大元素？

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
	递推公式：
		merge_sort(l…r) = merge(merge_sort(l…m), merge_sort(m+1…r))
	终止条件：
		l >= r 不用再继续分解
	归并排序是稳定的排序算法吗？
		稳定排序算法
	归并排序的时间复杂度是多少？
		不仅递归求解的问题可以写成递推公式，递归代码的时间复杂度也可以写成递推公式
		利用递归公式，求出归并排序的时间复杂度
			T(1) = C：n=1时，只需要常量级的执行时间，所以表示为C
			T(n) = 2*T(n/2) + n：n>1
		求解 T(n) = 2*T(n/2) + n
			 = 2*(2*T(n/4) + n/2) + n = 4*T(n/4) + 2*n
			 = 4*(2*T(n/8) + n/4) + 2*n = 8*T(n/8) + 3*n
			 = 8*(2*T(n/16) + n/8) + 3*n = 16*T(n/16) + 4*n
			 ......
			 = 2^k * T(n/2^k) + k * n
			 ......
		O(nlogn)
			当 T(n/2^k)=T(1) 时，也就是 n/2^k=1，我们得到 k=log2n
			将 k 值代入上面的公式，得到 T(n)=Cn+nlog2n
		归并排序的执行效率与要排序的原始数组的有序程度无关
			所以其时间复杂度是非常稳定的，不管是最好情况、最坏情况，还是平均情况，时间复杂度都是 O(nlogn)
	归并排序的空间复杂度是多少？
		1.不是原地排序算法
		2.通过递推公式来求解，那整个归并过程需要的空间复杂度就是 O(nlogn)
		3.在任意时刻，CPU 只会有一个函数在执行，也就只会有一个临时的内存空间在使用
			临时内存空间最大也不会超过 n 个数据的大小，所以空间复杂度是 O(n)

快速排序原理（Quicksort）
	分治思想
原理
	分区点 pivot

快速排序的性能分析
	递推公式：
		quick_sort(l…r) = quick_sort(l…pivot-1) + quick_sort(pivot+1…r)
	终止条件：
		l >= r
	快速排序是稳定的排序算法吗？
		非稳定排序算法
	快速排序的时间复杂度是多少？
		利用递归公式，求出归并排序的时间复杂度
			T(1) = C：n=1时，只需要常量级的执行时间，所以表示为C
			T(n) = 2*T(n/2) + n：n>1
		最好：O(nlogn)，分区极其均衡
		最坏：O(n^2)，分区极端不均衡
			比如选择区间的最后一个元素作为 pivot，当数组原来已经有序
			则需要进行大约 n 次分区，每次分区平均要扫描大约 n/2 个元素
		平均：O(nlogn)
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
	利用哨兵简化归并排序的 merge 函数
		E:\gothmslee\algo\sort\2.4MergeSort.go
	现在你有 10 个接口访问日志文件，每个日志文件大小约 300MB，每个文件里的日志都是按照时间戳从小到大排序的
	你希望将这 10 个较小的日志文件，合并为 1 个日志文件，合并之后的日志仍然按照时间戳从小到大排列
	如果处理上述排序任务的机器内存只有 1GB，你有什么好的解决思路，能“快速”地将这 10 个日志文件合并吗？
	参考
		先构建十条io流，分别指向十个文件，每条io流读取对应文件的第一条数据
		然后比较时间戳，选择出时间戳最小的那条数据，将其写入一个新的文件，然后指向该时间戳的io流读取下一行数据，然后继续刚才的操作，比较选出最小的时间戳数据，写入新文件，io流读取下一行数据
		以此类推，完成文件的合并
		这种处理方式，日志文件有n个数据就要比较n次，每次比较选出一条数据来写入，时间复杂度是O（n），空间复杂度是O（1）,几乎不占用内存
	GPT
		将每个日志文件划分为块：
			将每个日志文件划分为可以放入内存的较小块。由于每个文件大小约为 300MB，而我们有 1GB 的内存，因此我们可以选择每个块大小约为 100MB。
		对每个块进行排序：
			将每个块加载到内存中，根据时间戳进行排序，并将排序后的块写回磁盘。我们可以使用类似快速排序或归并排序的内存排序算法来完成此步骤。
		合并排序后的块：
			对所有 10 个文件的排序块执行 k 路归并。我们使用优先队列来选择所有块中下一个最小的时间戳条目。随着我们合并条目，我们将它们写入输出文件。
			由于我们一次只加载每个文件的一小部分到内存中，因此这一步可以高效地完成。
		通过将排序和合并过程分解为较小的可管理块，我们可以在内存限制下高效地将 10 个日志文件合并为一个排序的文件。
	个人
		桶排序+归并排序（+优先队列）
		桶排序比较适合用在外部排序中。数据存储在外部磁盘中，数据量比较大，内存有限，无法将数据全部加载到内存中
		思路同 GPT
*/
