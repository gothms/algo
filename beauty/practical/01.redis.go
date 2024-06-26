package practical

/*
算法实战（一）：剖析Redis常用数据类型对应的数据结构

经典数据库 Redis 中的常用数据类型，底层都是用哪种数据结构实现的？

Redis 数据库介绍
	语言
		C语言实现
	非关系型数据库
		Redis 是一种键值（Key-Value）数据库
		相对于关系型数据库（比如 MySQL），Redis 也被叫作非关系型数据库
	读写效率
		Redis 中只包含“键”和“值”两部分，只能通过“键”来查询“值”
		正是因为这样简单的存储结构，也让 Redis 的读写效率非常高
	内存数据库
		数据是存储在内存中的
		但是，它也支持将数据存储在硬盘中
	键值类型
		键的数据类型是字符串
		值的数据类型有很多，常用的数据类型有字符串、列表、字典、集合、有序集合

列表（list）
	两种实现方式
		压缩列表（ziplist）：存储的数据量比较小，需要同时满足下面两个条件
			列表中保存的单个数据（有可能是字符串类型的）小于 64 字节
			列表中数据个数少于 512 个
		双向循环链表：存储的数据量比较大
			当列表中存储的数据量比较大的时候，也就是不能同时满足上面两个条件的时候，列表就要通过双向循环链表来实现了
	压缩列表
		不是基础数据结构，而是 Redis 自己设计的一种数据存储结构
		有点儿类似数组，通过一片连续的内存空间，来存储数据
		跟数组不同的一点是，它允许存储的数据大小不同
		示例：01.redis_ziplist.jpg
	vs
		数组
			1.要求每个元素的大小相同，如果要存储不同长度的字符串，那就需要用最大长度的字符串大小作为元素的大小（假设是 20 个字节）
				当存储小于 20 个字节长度的字符串的时候，便会浪费部分存储空间
			2.同类型数据的存储
			3.随机访问
		压缩列表：特点
			1.比较节省内存（也是连续的内存空间）
			2.可以支持不同类型数据的存储
			3.因为数据存储在一片连续的内存空间，通过键来获取值为列表类型的数据，读取的效率也非常高
			4.不支持随机访问（因为通过 key 获取 value，所以不需要随机访问）
	Redis.list 的实现
		额外定义一个 list 结构体，来组织链表的首、尾指针，还有长度等信息
		typedef struct listnode {
		  struct listNode *prev;
		  struct listNode *next;
		  void *value;
		} listNode;
		typedef struct list {
		  listNode *head;
		  listNode *tail;
		  unsigned long len;
		  // ....省略其他定义
		} list;
字典（hash）
	两种实现方式
		压缩列表（ziplist）：存储的数据量比较小，需要同时满足下面两个条件
			字典中保存的键和值的大小都要小于 64 字节
			字典中键值对的个数要小于 512 个
		散列表：存储的数据量比较大
			当不能同时满足上面两个条件的时候，Redis 就使用散列表来实现字典类型
	hash
		hash算法
			Redis 使用 MurmurHash2 作为哈希函数
			https://zh.wikipedia.org/wiki/Murmur%E5%93%88%E5%B8%8C
		MurmurHash2 特点
			运行快
			随机性好
		哈希冲突
			链表法
		动态扩容
			装载因子大于 1 时，将散列表扩大为原来大小的 2 倍左右（具体值需要计算才能得到，可以阅读源码）
			源码：https://github.com/redis/redis/blob/unstable/src/dict.c
		动态缩容
			装载因子小于 0.1 的时候，Redis 就会触发缩容，缩小为字典中数据个数的大约 2 倍大小（也是计算得到，可以阅读源码）
			源码：https://github.com/redis/redis/blob/unstable/src/dict.c
		渐进式扩容缩容策略
			扩容缩容要做大量的数据搬移和哈希值的重新计算，所以比较耗时
			Redis 使用渐进式扩容缩容策略，将数据的搬移分批进行，避免了大量数据一次性搬移导致的服务停顿
			参见 11.hash_table_02.go
集合（set）
	数据特点
		存储一组不重复的数据
	两种实现方式
		有序数组：需要同时满足下面两个条件
			存储的数据都是整数
			存储的数据元素个数不超过 512 个
		散列表
			当不能同时满足这两个条件的时候，Redis 就使用散列表来存储集合中的数据
有序集合（sortedset）
	数据特点
		存储一组数据，并且每个数据会附带一个得分
		通过得分的大小，将数据组织成跳表这样的数据结构，以支持快速地按照得分值、得分区间获取数据
	两种实现方式
		压缩列表：存储的数据量比较小，，需要同时满足下面两个条件
			所有数据的大小都要小于 64 字节
			元素个数要小于 128 个
		跳表

数据结构持久化
	支持硬盘存储
		将内存中的数据存储到硬盘中
		当机器断电的时候，存储在 Redis 中的数据也不会丢失
		在机器重新启动之后，Redis 只需要再将存储在硬盘中的数据，重新读取到内存，就可以继续工作了
	如何将数据结构持久化到硬盘？
		问题描述
			像字典、集合等类型，底层用到了散列表，散列表中有指针的概念，而指针指向的是内存中的存储地址
			Redis 是如何将这样一个跟具体内存地址有关的数据结构存储到磁盘中的呢？
		数据结构的持久化问题
			Redis 遇到的这个问题并不特殊，很多场景中都会遇到
			叫做数据结构的持久化问题，或者对象的持久化问题
			“持久化”，可以笼统地理解为“存储到磁盘”
	两种解决思路：如何将数据结构持久化到硬盘？
		1.清除原有的存储结构，只将数据存储到磁盘中
			当需要从磁盘还原数据到内存的时候，再重新将数据组织成原来的数据结构
			Redis 采用的就是这种持久化思路
			弊端：
				数据从硬盘还原到内存的过程，会耗用比较多的时间
			举例：要将散列表中的数据存储到磁盘
				当从磁盘中，取出数据重新构建散列表的时候，需要重新计算每个数据的哈希值
				如果磁盘中存储的是几 GB 的数据，那重构数据结构的耗时就不可忽视了
		2.保留原来的存储格式，将数据按照原有的格式存储在磁盘中
			举例：要将散列表中的数据存储到磁盘
				将散列表的大小、每个数据被散列到的槽的编号等信息，都保存在磁盘中
				从磁盘中将数据还原到内存中的时候，就可以避免重新计算哈希值

Redis 中常用数据类型底层依赖的数据结构
	压缩列表（一种特殊的数组）
	有序数组
	链表
	散列表
	跳表

思考
	1.在数据量比较小的情况下，Redis 中的很多数据类型，比如字典、有序集合等，都是通过多种数据结构来实现的，为什么会这样设计呢？
	用一种固定的数据结构来实现，不是更加简单吗？
		参考：
		主要是出于时间和空间的考虑
		补充：
		1)压缩列表不支持随机访问，有点类似链表，但是比较省存储空间
		当想要随机访问某个元素的时候还是要像列表那样从头开始遍历，所以不能太大
		压缩列表的好处是能利用l2缓存，越小越有利于CPU缓存
		2)跳表相对于B+树更灵活，更容易实现，所以 redis 选择跳表而不是B+
		3)redis 用“清除格式，持久化数据再组织数据结构”，相对更消耗性能
		对于Redis来说，重启并不是很经常的事情。所以并不会经常从硬盘加载数据到内存再重构成数据结构
	2.数据结构持久化有两种方法。对于二叉查找树这种数据结构，如何将它持久化到磁盘中呢？
		参考：E:\gothmslee\algo\beauty\basic\13.tree_01.go
		对应方式一：中序遍历为数组后存储？
		对应方式二：“层序遍历”为数组后存储？可保有原有的二叉查找树的结构
*/
