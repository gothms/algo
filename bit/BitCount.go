package bit

// GetBitCount 获取位数
// Go 源码： bits.Len
func GetBitCount(i int) (n int) {
	if i>>16 != 0 {
		n += 16
		i >>= 16
	}
	if i>>8 != 0 {
		n += 8
		i >>= 8
	}
	if i>>4 != 0 {
		n += 4
		i >>= 4
	}
	if i>>2 != 0 {
		n += 2
		i >>= 2
	}
	if i>>1 != 0 {
		n += 1
		i >>= 1
	}
	return n + i
}

// GetLeadingZeros 获取前置 0 个数
// Go 源码： bits.LeadingZeros
func GetLeadingZeros(i int) (n int) {
	if i == 0 {
		return 32
	}
	if i>>16 == 0 {
		n += 16
		i <<= 16
	}
	if i>>24 == 0 {
		n += 8
		i <<= 8
	}
	if i>>28 == 0 {
		n += 4
		i <<= 4
	}
	if i>>30 == 0 {
		n += 2
		i <<= 2
	}
	n -= i>>31 - 1
	return n
}
