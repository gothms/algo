package advance

/*
索引：如何在海量数据中快速查找某个数据？

MySQL 底层依赖的是 B+ 树这种数据结构，那类似 Redis 这样的 Key-Value 数据库中的索引，又是怎么实现的呢？
底层依赖的又是什么数据结构呢？
索引这种常用的技术解决思路，底层往往会依赖哪些数据结构？

为什么需要索引？
	概述
		在实际的软件开发中，业务纷繁复杂，功能千变万化，但是，万变不离其宗
		如果抛开这些业务和功能的外壳，其实它们的本质都可以抽象为“对数据的存储和计算”
		对应到数据结构和算法中，那“存储”需要的就是数据结构，“计算”需要的就是算法
	存储
		对于存储的需求，功能上无外乎增删改查
		一旦存储的数据很多，那性能就成了这些系统要关注的重点
		特别是在一些跟存储相关的基础系统（比如 MySQL 数据库、分布式文件系统等）、中间件（比如消息中间件 RocketMQ 等）中
		重点：如何节省存储空间、如何提高数据增删改查的执行效率
			而这些系统的实现，都离不开一个东西，那就是索引
	索引
		索引设计得好坏，直接决定了这些系统是否优秀
		类比书籍的目录来理解。如果没有目录，我们想要查找某个知识点的时候，就要一页一页翻
		通过目录，就可以快速定位相关知识点的页数，查找的速度也会有质的提高

索引的需求定义
	对于系统设计需求，一般可以从功能性需求和非功能性需求两方面来分析
	1.功能性需求
	2.非功能性需求
功能性需求
	数据是格式化数据还是非格式化数据？
		要构建索引的原始数据，类型有很多，分为两类：
		1.结构化数据，比如，MySQL 中的数据
		2.非结构化数据，比如搜索引擎中网页
			一般需要做预处理，提取出查询关键词，对关键词构建索引
	数据是静态数据还是动态数据？
		1.静态数据
			不会有数据的增加、删除、更新操作
			只需要考虑查询效率就可以
		2.动态数据
			不仅要考虑到索引的查询效率，在原始数据更新的同时，还需要动态地更新索引
			支持动态数据集合的索引，设计起来相对也要更加复杂些
	索引存储在内存还是硬盘？
		1.存储在内存中
			查询的速度肯定要比存储在磁盘中的高
		2.存储在磁盘中
			如果原始数据量很大的情况下，对应的索引可能也会很大，因为内存有限，可能就不得不将索引存储在磁盘中了
		3.一部分存储在内存，一部分存储在磁盘
			兼顾内存消耗和查询效率
	单值查找还是区间查找？实际上，不同的应用场景，查询的需求会多种多样
		1.单值查找
			最常见的查询需求，根据查询关键词等于某个值的数据
		2.区间查找
			查找关键词处于某个区间值的所有数据
	单关键词查找还是多关键词组合查找？
		搜索引擎中构建的索引，既要支持一个关键词的查找，比如“数据结构”，也要支持组合关键词查找，比如“数据结构 AND 算法”
		1.单关键词查找
			索引构建起来相对简单些
		2.多关键词组合查找
			分多钟情况
			a)像 MySQL 这种结构化数据的查询需求，可以实现针对多个关键词的组合，建立索引
			b)对于像搜索引擎这样的非结构数据的查询需求，可以针对单个关键词构建索引
				然后通过集合操作，比如求并集、求交集等，计算出多个关键词组合的查询结果
	实际上，不同的场景，不同的原始数据，对于索引的需求也会千差万别
		以上值列举了一些比较有共性的需求
非功能性需求
	不管是存储在内存中还是磁盘中，索引对存储空间的消耗不能过大
		1.存储在内存中
			索引对占用存储空间的限制就会非常苛刻
			毕竟内存空间非常有限，一个中间件启动后就占用几个 GB 的内存，开发者显然是无法接受的
		2.存储在磁盘中
			索引对占用存储空间的限制，稍微会放宽一些
			也不能掉以轻心，因为有时候，索引对存储空间的消耗会超过原始数据
	在考虑索引查询效率的同时，还要考虑索引的维护成本
		索引的目的是提高查询效率
			但是，基于动态数据集合构建的索引，还要考虑到，索引的维护成本
		因为在原始数据动态增删改的同时，也需要动态地更新索引
			而索引的更新势必会影响到增删改操作的性能

构建索引常用的数据结构有哪些？
	动态数据结构
		散列表
		红黑树
		跳表
		B+树
	辅助索引
		位图
		布隆过滤器
	静态数据构建索引
		有序数组

	散列表
		键值数据库，比如 Redis、Memcache，就是使用散列表来构建索引的
		索引一般都构建在内存中
	红黑树
		非常适合用来构建内存索引
		Ext 文件系统中，对磁盘块的索引
	B+树
		比起红黑树来说，更加适合构建存储在磁盘中的索引
			B+ 树是一个多叉树，所以，对相同个数的数据构建索引，B+ 树的高度要低于红黑树
			当借助索引查询数据的时候，读取 B+ 树索引，需要的磁盘 IO 次数会更少
		大部分关系型数据库的索引，比如 MySQL、Oracle
	跳表
		过灵活调整索引结点个数和数据个数之间的比例，可以很好地平衡索引对内存的消耗及其查询效率
		Redis 中的有序集合
	位图 & 布隆过滤器
		用于索引中，辅助存储在磁盘中的索引，加速数据查找的效率
	布隆过滤器
		缺点：误判
		优点：内存占用非常少
		辅助：规避它的短处，发挥它的长处
			针对数据，构建一个布隆过滤器，并且存储在内存中
			要查询数据的时候，可以先通过布隆过滤器，判定是否存在
				如果通过布隆过滤器判定数据不存在，那我们就没有必要读取磁盘中的索引了
				对于数据不存在的情况，数据查询就更加快速了
	有序数组
		如果数据是静态的，也就是不会有插入、删除、更新操作
		可以把数据的关键词（查询用的）抽取出来，组织成有序数组，然后利用二分查找算法来快速查找数据

总结
	架构设计离不开数据结构和算法
	一个优秀的业务架构师、基础架构师，数据结构和算法的根基一定要打稳
	那些看似很惊艳的架构设计思路，实际上，都是来自最常用的数据结构和算法

思考
	基础系统、中间件、开源软件等系统中，有哪些用到了索引吗？这些系统的索引是如何实现的呢？
	// TODO
*/
