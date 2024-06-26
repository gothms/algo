package basic

/*
散列表（上）：Word文档中的单词拼写检查功能是如何实现的？

散列表（Hash Table）
	散列函数（或“Hash 函数”“哈希函数”）：把原数据转化为数组下标的映射方法
	散列值（或“Hash 值”“哈希值”）：散列函数计算得到的值

散列思想
	散列表用的是数组支持按照下标随机访问数据的特性，所以散列表其实就是数组的一种扩展，由数组演化而来
	可以说，如果没有数组，就没有散列表

	散列表用的就是数组支持按照下标随机访问的时候，时间复杂度是 O(1) 的特性
	通过散列函数把元素的键值映射为下标，然后将数据存储在数组中对应下标的位置
	当按照键值查询元素时，用同样的散列函数，将键值转化数组下标，从对应的数组下标的位置取数据

散列函数设计的基本要求
	散列函数计算得到的散列值是一个非负整数
	如果 key1 = key2，那 hash(key1) == hash(key2)
	如果 key1 ≠ key2，那 hash(key1) ≠ hash(key2)：著名的 MD5、SHA、CRC 等哈希算法，也无法完全避免这种散列冲突
散列冲突
	开放寻址法（open addressing）
	链表法（chaining）
开放寻址法
	线性探测（Linear Probing）
		插入：
			如果某个数据经过散列函数散列之后，存储位置已经被占用了，我们就从当前位置开始，依次往后查找，看是否有空闲位置，直到找到为止
		查找：
			通过散列函数求出要查找元素的键值对应的散列值，然后比较数组中下标为散列值的元素和要查找的元素
			如果相等，则说明就是我们要找的元素
			否则就顺序往后依次查找
			如果遍历到数组中的空闲位置，还没有找到，就说明要查找的元素并没有在散列表中
		删除：
			将删除的元素，特殊标记为 deleted（防止查找时该位置已删除，出现脏数据）
			当线性探测查找的时候，遇到标记为 deleted 的空间，并不是停下来，而是继续往下探测
		问题：
			当散列表中插入的数据越来越多时，散列冲突发生的可能性就会越来越大，空闲位置会越来越少，线性探测的时间就会越来越久
			极端情况下，我们可能需要探测整个散列表，所以最坏情况下的时间复杂度为 O(n)
			查找和删除也是同理
	二次探测（Quadratic probing）
		或叫平方探测法
		线性探测每次探测的步长是 1
		二次探测探测的步长就变成了原来的“二次方”，1 4 9 ...
	双重散列（Double hashing）
		先用第一个散列函数，如果计算得到的存储位置已经被占用，再用第二个散列函数，依次类推，直到找到空闲的存储位置
链表法
	在散列表中，每个“桶（bucket）”或者“槽（slot）”会对应一条链表，所有散列值相同的元素我们都放到相同槽位对应的链表中
	插入：O(1)
	查找 & 删除：跟链表的长度 k 成正比，也就是 O(k)
		对于散列比较均匀的散列函数来说，理论上讲，k=n/m，其中 n 表示散列中数据的个数，m 表示散列表中“槽”的个数

装载因子（load factor）
	散列表的装载因子=填入表中的元素个数/散列表的长度
	装载因子越大，说明空闲位置越少，冲突越多，散列表的性能会下降

Word 的这个单词拼写检查功能，虽然很小但却非常实用。你有没有想过，这个功能是如何实现的呢？
	英文单词有 20 万个左右，假设单词的平均长度是 10 个字母，平均一个单词占用 10 个字节的内存空间
	那 20 万英文单词大约占 2MB 的存储空间，就算放大 10 倍也就是 20MB，这个大小完全可以放在内存里面
	用散列表来存储整个英文单词词典
	当用户输入某个英文单词时，拿用户输入的单词去散列表中查找，可以轻松实现快速判断输入的单词拼写正确与否

小结
	散列表两个核心问题是散列函数设计和散列冲突解决
	散列冲突有两种常用的解决方法，开放寻址法和链表法
	散列函数设计的好坏决定了散列冲突的概率，也就决定散列表的性能

思考
	假设我们有 10 万条 URL 访问日志，如何按照访问次数给 URL 排序？
		遍历 10 万条数据，以 URL 为 key，访问次数为 value，存入散列表，同时记录下访问次数的最大值 K，时间复杂度 O(N)
		如果 K 不是很大，可以使用桶排序，时间复杂度 O(N)。如果 K 非常大（比如大于 10 万），就使用快速排序，复杂度 O(NlogN)
	有两个字符串数组，每个数组大约有 10 万条字符串，如何快速找出两个数组中相同的字符串？
		散列表
		k-v：字符串-出现次数
*/
