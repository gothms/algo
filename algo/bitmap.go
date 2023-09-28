package algo

// BitMap 位图
type BitMap struct {
	bytes []uint16
	n     int
}

func (b *BitMap) set(k int) {
	if k > b.n {
		return
	}
	idx, i := k>>4, k&15
	b.bytes[idx] |= 1 << i
}
func (b *BitMap) get(k int) bool {
	if k > b.n {
		return false
	}
	idx, i := k>>4, k&15
	return b.bytes[idx]&(1<<i) != 0 // 需要括号
}
func (b *BitMap) remove(k int) {
	if k > b.n {
		return
	}
	idx, i := k>>4, k&15
	b.bytes[idx] &^= 1 << i
}
func ibit(key int) (int, int) {
	return key >> 4, key & 15
}

// BloomFilter Google 的 Guava 工具包提供了 BloomFilter 布隆过滤器的实现
