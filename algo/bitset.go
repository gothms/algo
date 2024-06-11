package algo

import (
	"math/bits"
	"slices"
	"sort"
)

/*
bitset 包
	https://github.com/bits-and-blooms/bitset
	https://pkg.go.dev/github.com/bits-and-blooms/bitset#section-readme
*/

// ====================BitMap 原始版====================

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
func iBit(key int) (int, int) {
	return key >> 4, key & 15
}

// ====================BitSet 拷贝后再左移====================
// https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/solutions/2805413/bitset-you-hua-0-1-bei-bao-by-endlessche-m1xn/

const w = bits.UintSize

type bitSet []uint

// b <<= k
func (b bitSet) lsh(k int) bitSet {
	shift, offset := k/w, k%w
	if offset == 0 { // Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(w-offset)
		}
		b[shift] = b[0] << offset
	}
	clear(b[:shift])
	return b
}

// 把 >= start 的清零
func (b bitSet) resetRange(start int) bitSet {
	i := start / w
	b[i] &= uint(1)<<(start%w) - 1
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitSet) unionFrom(c bitSet) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitSet) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

func maxTotalReward_(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重
	m := rewardValues[len(rewardValues)-1]
	f := make(bitSet, (m*2+1)/w+1)
	f[0] = 1
	for _, val := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(val).resetRange(val * 2))
	}
	return f.lastIndex1()
}

// ====================BitSet 原地左移====================

// bitset 周赛 401
type bitset struct {
	bs []uint
	us int
}

// left TODO 画图
func (b bitset) left(k int) {
	shift, offset := k/b.us, k%b.us
	// 第一段：bitset 的起始 [0,l1] l1=us-offset 位，移到 shift.[offset,us)
	l1 := min(k, b.us-offset)
	b.bs[shift] = b.bs[shift] | (b.bs[0] & (1<<l1 - 1) << offset)
	if shift == 0 && offset < b.us>>1 { // 总共需移动 bitset[0] 的前 offset 位，且第一段已完成所有移动
		return
	}

	// 中间段：拼接两段被截断的 uint 为一整段，然后整段移到 bitset[shift+i+1]
	for i := 0; i < (k-b.us+offset)/b.us; i++ {
		b.bs[shift+i+1] = b.bs[shift+i+1] | (b.bs[i]>>l1 | b.bs[i+1]<<offset)
	}

	// 第三段：整段移位后剩余的位，移到末尾
	l3 := (offset << 1) % b.us // 第三段移动 l3 位，offset 为 bs[shift] 在 k 截断后的“前缀”
	if l3 <= offset {          // “前缀”足够，不需要借位
		if shift<<1+1 >= len(b.bs) { // shift<<1+1 可能刚好越界
			return
		}
		// 截取 bitset[shift] 的 offset 前缀的 l3 长度位，移动到 bitset[shift<<1+1] 的前 l3 位
		b.bs[shift<<1+1] = b.bs[shift<<1+1] | b.bs[shift]>>(offset-l3)&(1<<l3-1)
	} else {
		// 向 shift-1 借 l3-offset 位，然后拼到低位，最后移动到 bitset[shift<<1]
		b.bs[shift<<1] |= b.bs[shift]&(1<<offset-1)<<(l3-offset) | (b.bs[shift-1] >> (b.us - l3 + offset))
	}
}

// length 获取 bitset 的最大数
func (b bitset) length() int {
	for i := len(b.bs) - 1; i >= 0; i-- {
		if b.bs[i] != 0 {
			return i*b.us | (bits.Len(b.bs[i]) - 1)
		}
	}
	return -1
}
func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	us := bits.UintSize
	bs := make([]uint, (rewardValues[len(rewardValues)-1]<<1+1)/us+1)
	bs[0] = 1
	bitSet := bitset{bs, us}
	for _, v := range rewardValues {
		bitSet.left(v)
	}
	return bitSet.length()
}

// BloomFilter Google 的 Guava 工具包提供了 BloomFilter 布隆过滤器的实现
