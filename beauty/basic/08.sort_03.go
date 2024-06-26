package basic

/*
线性排序：如何根据年龄给100万用户数据排序？

线性排序（Linear sort）
	线性的时间复杂度
	非基于比较的排序算法，不涉及元素之间的比较操作

桶排序（Bucket sort）
	核心思想是将要排序的数据分到几个有序的桶里，每个桶里的数据再单独进行排序
	桶内排完序之后，再把每个桶里的数据按照顺序依次取出，组成的序列就是有序的了
时间复杂度
	如果要排序的数据有 n 个，我们把它们均匀地划分到 m 个桶内，每个桶里就有 k=n/m 个元素
	每个桶内部使用快速排序，时间复杂度为 O(k * logk)
	m 个桶排序的时间复杂度就是 O(m * k * logk)，因为 k=n/m，所以整个桶排序的时间复杂度就是 O(n*log(n/m))
	当桶的个数 m 接近数据个数 n 时，log(n/m) 就是一个非常小的常量，这个时候桶排序的时间复杂度接近 O(n)
适用条件
	要排序的数据需要很容易就能划分成 m 个桶，并且，桶与桶之间有着天然的大小顺序
	数据在各个桶之间的分布是比较均匀的
	比较适合用在外部排序中
		外部排序就是数据存储在外部磁盘中，数据量比较大，内存有限，无法将数据全部加载到内存中
		某些桶的数据特别多，可以继续划分
示例
	有 10GB 的订单数据，我们希望按订单金额（假设金额都是正整数）进行排序，但是内存有限，只有几百 MB，没办法一次性把 10GB 的数据都加载到内存中。这个时候该怎么办呢？
	实现步骤：
		可以先扫描一遍文件，看订单金额所处的数据范围
		将所有订单根据金额划分到 100 个桶里，第一个桶我们存储金额在 1 元到 1000 元之内的订单
		第二桶存储金额在 1001 元到 2000 元之内的订单，以此类推
		每一个桶对应一个文件，并且按照金额范围的大小顺序编号命名（00，01，02...99）
		每个小文件中存储大约 100MB 的订单数据，就可以将这 100 个小文件依次放到内存中，用快排来排序
		按照文件编号，从小到大依次读取每个小文件中的订单数据，并将其写入到一个文件中
	细节：
		订单按照金额在 1 元到 10 万元之间并不一定是均匀分布的
		针对这些划分之后还是比较大的文件，可以继续划分
		直到所有的文件都能读入内存为止
	参见：23.divide_conquer.go

计数排序（Counting sort）
	计数排序其实是桶排序的一种特殊情况
	稳定排序写法
	非稳定排序写法
适用条件
	只能用在数据范围不大的场景中，如果数据范围 k 比要排序的数据 n 大很多，就不适合用计数排序了
	而且，计数排序只能给非负整数排序，如果要排序的数据是其他类型的，要将其在不改变相对大小的情况下，转化为非负整数

基数排序（Radix sort）：思路类似使用稳定排序算法来排序订单的金额、下单时间
	数据可以划分成高低位，位之间有递进关系
	按照每位来排序的排序算法要是稳定的
	根据每一位来排序，可以用桶排序或者计数排序，它们的时间复杂度可以做到 O(n)
示例：不等长排序，排序牛津字典中的 20 万个英文单词
	把所有的单词补齐到相同长度，位数不够的可以在后面补“0”
	再用基数排序进行排序
适用条件
	要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系
	如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了
	除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序
	否则，基数排序的时间复杂度就无法做到 O(n) 了

小结
	桶排序和计数排序的排序思想是非常相似的，都是针对范围不大的数据，将数据划分成不同的桶来实现排序
	基数排序要求数据可以划分成高低位，位之间有递进关系
		比较两个数，只需要比较高位，高位相同的再比较低位
		而且每一位的数据范围不能太大，因为基数排序算法需要借助桶排序或者计数排序来完成每一个位的排序工作
	三者都是稳定排序算法（或可实现）
	三者都不是原地排序算法
如何根据年龄给 100 万用户排序？
	计数排序

思考
	假设我们现在需要对 D，a，F，B，c，A，z 这个字符串进行排序，要求将其中所有小写字母都排在大写字母的前面
	但小写字母内部和大写字母内部不要求有序。比如经过排序之后为 a，c，z，D，F，B，A，这个如何来实现呢？
		双指针：l r，l 遇到大写字母停下，r 遇到小写字母停下，交换 l r
	如果字符串中存储的不仅有大小写字母，还有数字
	要将小写字母的放到前面，大写字母放在最后，数字放在中间，不用排序算法，又该怎么解决呢？
		1.双指针：两次，第一次将数字看做小写
		2.将三种数据放入三个集合中，再汇总这三个集合
	假设有 1 亿个整数，数据范围是从 1 到 10 亿，如何快速并且省内存地给这 1 亿个数据从小到大排序？
		参见 advance.03.bitmap_bloom-filter.go 思考 1.

		GPT
			对于如此大规模的数据，我们需要一种快速且省内存的排序算法。针对这种情况，外部排序算法是最合适的选择。外部排序算法通常用于排序大规模数据，因为它们能够处理超出内存容量的数据集。

			下面是一种可能的方法，使用外部排序算法来排序这1亿个整数：

			分割数据为小块: 将1亿个整数分割成适当大小的块，每个块的大小适合内存加载并进行排序。比如将1亿个数分为1000个块，每个块大约包含10万个数。
			内部排序: 将每个块加载到内存中进行排序。由于每个块的大小适合内存，我们可以使用快速排序、归并排序等高效的内部排序算法来排序这些块。
			归并排序: 使用归并排序算法对排序好的块进行归并，以得到最终的排序结果。在归并过程中，我们可以利用外部存储来存储中间结果，而不是一次性加载整个数据集到内存中。
			这样，通过分块、内部排序和归并排序，我们可以在不超出内存容量的情况下，快速地对1亿个整数进行排序。

			在实际操作中，我们可以选择使用外部存储设备（如硬盘、SSD）来存储中间结果和块数据。另外，对于内部排序算法和归并排序算法的选择，需要考虑到具体情况，以确保排序的效率和内存使用的合理性。
		个人
			基数排序
			先将数字按位数区分，对于 1位、2位 等包含数据较少的区间，可以采用快速/归并/插入排序等
			而对于位数较大的区间，如 [一亿,十亿] 的数据，采用基数排序
		https://blog.csdn.net/qq_39521554/article/details/79546854
		// TODO
*/
