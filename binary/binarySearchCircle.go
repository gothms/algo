package binary

// BinarySearchCircle 查找 旋转 有序数组
func BinarySearchCircle(arr []int, t int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if arr[mid] == t {
			return mid
		} else if arr[0] > arr[mid] {
			if arr[mid] < t && arr[r] >= t {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if arr[mid] > t && arr[l] <= t {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
