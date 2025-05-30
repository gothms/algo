package _1_basic

/*
如何使用数据结构表达树？
	数组适合快速地随机访问。不过，数组并不适合稀疏的数列或者矩阵，而且数组中元素的插入和删除操作也比较低效
	相对于数组，链表的随机访问的效率更低，但是它的优势是，不必事先规定数据的数量，表示稀疏的数列或矩阵时，可以更有效地利用存储空间，同时也利于数据的动态插入和删除

	树的结点及其之间的边，和链表中的结点和链接在本质上是一样的，因此，我们可以模仿链表的结构，用编程语言中的指针或对象引用来构建树
	除此之外，我们其实还可以用二维数组。用数组的行或列元素表示树中的结点，而行和列共同确定了两个树结点之间是不是存在边
		可是在树中，这种二维关系通常是非常稀疏的、非常动态的，所以用数组效率就比较低下

	设计一个 TreeNode 类，表示有向树的结点和边
		这个结点所代表的字符，要用 label 变量表示
		这个结点有哪些子结点，要用 sons 哈希映射表示。之所以用哈希，是为了便于查找某个子结点（或者说对应的字符）是否存在
			实际上，这里的有向边表达的是父子结点之间的关系，我把这种关系用 sons 变量来存储子结点
		还可以用变量 prefix 表示当前结点之前的前缀，用变量 explanation 表示某个单词的解释

如何使用递归和栈实现深度优先搜索？
	什么样的编程方式可以实现对树结点和边的操作？
		深度优先搜索的过程和递归调用在逻辑上是一致的
			仔细观察前缀树构建和查询，你会发现这两个不断重复迭代的过程，都可以使用递归编程来实现
		栈
递归，在查询的过程中，至少有三种情况是无法在字典里找到被查的单词的
	第一种情况：被查单词所有字母都被处理完毕，但是我们仍然无法在字典里找到相应的词条
		确认一下当前匹配到的结点本身是不是一个单词。如果是，就返回相应的单词解释，否则就返回查找失败
	第二种情况：搜索到前缀树的叶子结点，但是被查单词仍有未处理的字母，就返回查找失败
	第三种情况：搜索到中途，还没到达叶子结点，被查单词也有尚未处理的字母，但是当前被处理的字母已经无法和结点上的 label 匹配，返回查找失败。是不是叶子仍然通过结点对象的 sons 变量来判断
如果我想遍历整个字典中所有的单词，那该怎么办呢？
	查找一个单词的过程，其实就是在有向树中，找一条从树的根到代表这个单词的结点之通路
		那么如果要遍历所有的单词，就意味着我们要找出从根到所有代表单词的结点之通路
	尽管函数递归调用非常直观，可是也有它自身的弱点
		函数的每次嵌套，都可能产生新的变量来保存中间结果，这可能会消耗大量的内存
		所以这里我们可以用一个更节省内存的数据结构，栈（Stack）
	为什么栈可以进行深度优先搜索呢？
		栈先进后出的特性，可以模拟函数的递归调用。实际上，计算机系统里的函数递归，在内部也是通过栈来实现的
		如果我们不使用函数调用时自动生成的栈，而是手动使用栈的数据结构，就能始终保持数据的副本只有一个，大大节省内存的使用量

归并排序、排列组合等课题，也采用了递归来实现，那它们是不是也算深度优先搜索呢？
	在归并排序的数据分解阶段，初始的数据集就是树的根结点，二分之前的数据集代表父节点，而二分之后的左半边的数据集和右半边的数据集都是父结点的子结点
		分解过程一直持续到单个的数值，也就是最末端的叶子结点，很明显这个阶段可以用树来表示
		如果使用递归编程来进行数据的切分，那么这种实现就是深度优先搜索的体现
	在排列中，我们可以把空集认为是树的根结点，如果把每次选择的元素作为父结点，那么剩下可选择的元素，就构成了这个父结点的子结点
		而每多选择一个元素，就会把树的高度加 1。因此，我们也可以使用递归和深度优先搜索，列举所有可能的排列

思考
	这两节我讲的是树的深度优先搜索。如果是在一般的图中进行深度优先搜索，会有什么不同呢？
		广度优先搜索可以使用同样的方式来遍历有多个连通子图的图
*/
