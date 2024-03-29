package binary

func Search(n int, f func(i int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

// BinarySearchFirstEqual 查找 第一个等于 给定值的元素
func BinarySearchFirstEqual(arr []int, t int) int {
	l, r, mid := 0, len(arr)-1, 0
	for l < r {
		mid = int(uint(l+r) >> 1)
		if arr[mid] < t {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if arr[l] == t {
		return l
	}
	return -1
}

// BinarySearchLastEqual 查找 最后一个等于 给定值的元素
func BinarySearchLastEqual(arr []int, t int) int {
	l, r, mid := 0, len(arr)-1, 0
	for l < r {
		mid = int(uint(l+r) >> 1)
		if arr[mid] <= t {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if arr[r] == t {
		return r
	}
	return -1
}

// BinarySearchFirstBigger 查找 第一个大于等于 给定值的元素
func BinarySearchFirstBigger(arr []int, t int) int {
	l, r, mid := 0, len(arr), 0
	for l < r {
		mid = int(uint(l+r) >> 1)
		if arr[mid] < t {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l // 数组为空时，l=0 越界
}

// BinarySearchLastSmaller 查找 最后一个小于等于 给定值的元素
func BinarySearchLastSmaller(arr []int, t int) int {
	l, r, mid := -1, len(arr)-1, 0
	for l < r { // 拦住了 r=-1
		mid = int(uint(l+r+1) >> 1)
		if arr[mid] <= t {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l // 错误：t<arr[0] 时返回了 0，则 l 初始化为 -1

	//l, r, mid := 0, len(arr), 0
	//for l < r {
	//	mid = int(uint(l+r) >> 1)
	//	if arr[mid] <= t {
	//		l = mid + 1
	//	} else {
	//		r = mid
	//	}
	//}
	//return r - 1
}

// BinarySearchLastSmallerGeekBang 查找 最后一个小于等于 给定值的元素
func BinarySearchLastSmallerGeekBang(arr []int, t int) int {
	l, r, n, mid := 0, len(arr)-1, len(arr)-1, 0
	for l <= r {
		mid = int(uint(l+r) >> 1)
		if arr[mid] > t {
			r = mid - 1
		} else {
			if mid == n || arr[mid+1] > t {
				return mid
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
