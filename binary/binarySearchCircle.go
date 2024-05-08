package binary

// BinarySearchCircle 查找 旋转 有序数组
func BinarySearchCircle(arr []int, t int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if arr[mid] == t {
			return mid
		} else if arr[0] > arr[mid] { // m 落在旋转数组的后半截
			if arr[mid] < t && arr[r] >= t { // 往后半截的 right 移
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else { // m 落在旋转数组的前半截
			if arr[mid] > t && arr[l] <= t { // 往前半截的 left 移
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

// 写法二
func BinarySearchCircle_(arr []int, t int) int {
	l, r, mid := 0, len(arr)-1, 0
	for l <= r {
		mid = (l + r) >> 1
		if arr[mid] == t {
			return mid
		} else if arr[mid] < t {
			if arr[0] > arr[mid] && arr[0] <= t {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if arr[0] > t && arr[0] <= arr[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
