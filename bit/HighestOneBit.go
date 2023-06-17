package bit

// GetHighestOneBit 获取最高位的 1
func GetHighestOneBit(c int) int {
	c = c | (c >> 1)
	c = c | (c >> 2)
	c = c | (c >> 4)
	c = c | (c >> 8)
	c = c | (c >> 16)
	return c>>1 + 1
}

// MinQuantity 大于等于 v 的最小 2的幂
func MinQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
