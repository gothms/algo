package sorts

/*
	这里展示：zsortinterface.go
	另外有：zsortfunc.go

sort.Slice()
	insertionSort_func：插入排序，长度小于等于 12
	heapSort_func：堆排序
	pdqsort_func：快排
	partition_func
	choosePivot_func：分区算法
		[0,8]：选择一个静态支点
		[8,shortestNinther)：使用简单的三者中位数方法，i j k
		[shortestNinther,∞)：使用Tukey ninther方法，i-1 i i+1 取中(j k 同理)，再3数取中

		shortestNinther = 50
		swaps int	// 和逆序度有关，如果 == 12，则先反转原数据，并重新计算 pivot
		l := b - a
		var (
			swaps int
			i     = a + l/4*1
			j     = a + l/4*2
			k     = a + l/4*3
		)
	symMerge 合并两个排序子序列
		E:\gothmslee\algo\sort\2.5MergeSortGo.go
*/

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
func stable(data Interface, n int) {
	blockSize := 20 // must be > 0
	a, b := 0, blockSize
	for b <= n {
		insertionSort(data, a, b) // 插入排序
		a = b
		b += blockSize
	}
	insertionSort(data, a, n)

	for blockSize < n { // blockSize *= 2
		a, b = 0, 2*blockSize
		for b <= n {
			symMerge(data, a, a+blockSize, b) // 两两合并
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			symMerge(data, a, m, n)
		}
		blockSize *= 2
	}
}

// symMerge 使用 Pok-Son Kim 和 Arne Kutzner 的 SymMerge 算法合并两个排序子序列 data[a:m] 和 data[m:b]，"Stable Minimum Storage Merging by Symmetric Comparisons"，
// 载于 Susanne Albers 和 Tomasz Radzik 编辑，《算法 - ESA 2004》，《计算机科学讲义》第 3221 卷，第 714-723 页。Springer, 2004.
//
// 设 M = m-a 和 N = b-n。Wolog M < N.
// 递归深度受 ceil(log(N+M)) 约束。
// 该算法需要 O(M*log(N/M + 1)) 次调用 data.Less。
// 该算法需要 O((M+N)*log(M)) 次调用 data.Swap。
//
// 该论文给出了 O((M+N)*log(M))的赋值次数，假设旋转算法使用 O(M+N)*log(M) 旋转算法需要 O(M+N+gcd(M+N)) 次赋值。
// 论文中的论证 论文中的论证也适用于交换操作，尤其是块交换旋转只需 O(M+N+gcd(M+N)) 尤其是块交换旋转只使用 O(M+N) 次交换。
//
// symMerge 假定参数不退化：a < m && m < b。
// 通过让调用者检查这一条件，可以消除许多叶递归调用、 从而提高性能。

// symMerge merges the two sorted subsequences data[a:m] and data[m:b] using
// the SymMerge algorithm from Pok-Son Kim and Arne Kutzner, "Stable Minimum
// Storage Merging by Symmetric Comparisons", in Susanne Albers and Tomasz
// Radzik, editors, Algorithms - ESA 2004, volume 3221 of Lecture Notes in
// Computer Science, pages 714-723. Springer, 2004.
//
// Let M = m-a and N = b-n. Wolog M < N.
// The recursion depth is bound by ceil(log(N+M)).
// The algorithm needs O(M*log(N/M + 1)) calls to data.Less.
// The algorithm needs O((M+N)*log(M)) calls to data.Swap.
//
// The paper gives O((M+N)*log(M)) as the number of assignments assuming a
// rotation algorithm which uses O(M+N+gcd(M+N)) assignments. The argumentation
// in the paper carries through for Swap operations, especially as the block
// swapping rotate uses only O(M+N) Swaps.
//
// symMerge assumes non-degenerate arguments: a < m && m < b.
// Having the caller check this condition eliminates many leaf recursion calls,
// which improves performance.
func symMerge(data Interface, a, m, b int) {
	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[a] into data[m:b]
	// if data[a:m] only contains one element.
	if m-a == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] >= data[a] for m <= i < b.
		// Exit the search loop with i == b in case no such index exists.
		i := m
		j := b
		for i < j {
			h := int(uint(i+j) >> 1)
			if data.Less(h, a) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[a] reaches the position before i.
		for k := a; k < i-1; k++ {
			data.Swap(k, k+1)
		}
		return
	}

	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[m] into data[a:m]
	// if data[m:b] only contains one element.
	if b-m == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] > data[m] for a <= i < m.
		// Exit the search loop with i == m in case no such index exists.
		i := a
		j := m
		for i < j {
			h := int(uint(i+j) >> 1)
			if !data.Less(m, h) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[m] reaches the position i.
		for k := m; k > i; k-- {
			data.Swap(k, k-1)
		}
		return
	}

	mid := int(uint(a+b) >> 1)
	n := mid + m // 修正
	var start, r int
	if m > mid {
		start = n - b
		r = mid
	} else {
		start = a
		r = m
	}
	p := n - 1
	// 核心逻辑
	// 开始循环前，[start,r-1]、[r,p] 分别为有序区间
	// 一般情况下 p-c 和 c 的关系：[c,r) 和 [r,p-c] 的元素个数相同
	// 如果 Less() 是 <=，则找出 data[p-c] <= data[c] 的起始索引 c
	for start < r {
		c := int(uint(start+r) >> 1)
		if !data.Less(p-c, c) {
			start = c + 1
		} else {
			r = c
		}
	}

	end := n - start
	if start < m && m < end {
		// [start,m-1] 有序区间最小值 >= [m,end) 有序区间最大值
		// 两个区间的大小可能不一样，但无论如何 rotate函数 都能将两个区间的元素“对调”
		// 对调后，[start, end] 区间是有序的
		rotate(data, start, m, end)
	}
	if a < start && start < mid {
		symMerge(data, a, start, mid)
	}
	if mid < end && end < b {
		symMerge(data, mid, end, b)
	}
}

// rotate rotates two consecutive blocks u = data[a:m] and v = data[m:b] in data:
// Data of the form 'x u v y' is changed to 'x v u y'.
// rotate performs at most b-a many calls to data.Swap,
// and it assumes non-degenerate arguments: a < m && m < b.
func rotate(data Interface, a, m, b int) {
	i := m - a
	j := b - m

	for i != j {
		if i > j {
			swapRange(data, m-i, m, j)
			i -= j
		} else {
			swapRange(data, m-i, m+j-i, i)
			j -= i
		}
	}
	// i == j
	swapRange(data, m-i, m, i)
}
func swapRange(data Interface, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a+i, b+i)
	}
}

// ==========================================================
// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
