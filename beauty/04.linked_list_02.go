package beauty

/*
写出正确的链表代码

	技巧一：理解指针或引用的含义
		不管是“指针”还是“引用”，它们实际上的意思都是一样的，都是存储所指对象的内存地址
		将某个变量赋值给指针，实际上就是将这个变量的地址赋值给指针
		或者反过来说，指针中存储了这个变量的内存地址，指向了这个变量，通过指针就能找到这个变量。
	技巧二：警惕指针丢失和内存泄漏
		插入结点时，一定要注意操作的顺序
		示例：
		p->next = x;  // 将p的next指针指向x结点；
		x->next = p->next;  // 将x的结点的next指针指向b结点；
		如果是 C 语言，内存管理是由程序员负责的，如果没有手动释放结点对应的内存空间，就会产生内存泄露
		同理，删除链表结点时，也一定要记得手动释放内存空间
	技巧三：利用哨兵简化实现难度
		哨兵，解决的是国家之间的边界问题。同理，这里说的哨兵也是解决“边界问题”的，不直接参与业务逻辑

		场景：
		空链表中插入第一个结点
			加入哨兵结点，成为带头链表
		删除链表中的最后一个结点
		插入排序
		归并排序
		动态规划
		...
	技巧四：重点留意边界条件处理
		常用来检查链表代码是否正确的边界条件有这样几个
		如果链表为空时，代码是否能正常工作？
		如果链表只包含一个结点时，代码是否能正常工作？
		如果链表只包含两个结点时，代码是否能正常工作？
		代码逻辑在处理头结点和尾结点的时候，是否能正常工作？
	技巧五：举例画图，辅助思考
		举例法
		画图法
	技巧六：多写多练，没有捷径
		单链表反转：lc-206
		链表中环的检测：lc-141
		两个有序的链表合并：lc-21
		删除链表倒数第 n 个结点：lc-19
		求链表的中间结点：lc-876

思考
	还有那些场景，利用哨兵可以大大地简化编码难度
*/

// find 在数组arr中，查询值t所在索引，没找到则返回-1
// 举例说明哨兵的作用，写代码的时候千万不要写这样的代码，因为可读性太差了
// 大部分情况下，我们并不需要如此追求极致的性能
func find(arr []int, t int) int {
	n := len(arr) - 1
	if arr[n] == t { // 先比较哨兵要占的位置
		return n
	}
	temp, i := arr[n], 0 // 保存 arr[len(arr)-1]
	arr[n] = t           // 设置哨兵
	for arr[i] != t {    // 少了 i<=n 的判断
		i++
	}
	arr[n] = temp // 恢复数据
	if i == n {
		return -1 // 没找到
	}
	return i
}