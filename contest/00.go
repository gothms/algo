package main

import (
	"fmt"
	"math/bits"
	"slices"
	"sort"
)

func main() {
	//n, k := 3, 7
	////n, k = 5, 6 // 2
	//child := numberOfChild(n, k)
	//fmt.Println(child)

	//n, k := 4, 5
	//seconds := valueAfterKSeconds(n, k)
	//fmt.Println(seconds)

	rewardValues := []int{1, 6, 4, 3, 2}                                   // 11
	rewardValues = []int{2, 15, 14, 18}                                    // 35
	rewardValues = []int{1, 1, 3, 3}                                       // 4
	rewardValues = []int{49, 413, 190}                                     // 652
	rewardValues = []int{39, 72, 38, 93, 74}                               // 170
	rewardValues = []int{63, 6, 73, 54, 86, 72, 20, 91, 64, 6, 91, 55, 75} // 181
	rewardValues = []int{65, 88, 32}                                       // 153
	rewardValues = []int{1, 2, 4, 8, 16, 32, 64}                           // 127
	reward := maxTotalReward(rewardValues)
	fmt.Println(reward)
	reward2 := maxTotalReward2(rewardValues)
	fmt.Println(reward2)

	//fmt.Println(bits.UintSize)
	var i int
	v := ^(^uint(0) << i)
	fmt.Printf("%b\n", v)
}
func numberOfChild(n int, k int) int {
	//if k == 0 {
	//	return 0
	//}
	//if k < n {
	//	return k
	//}
	i := k / (n - 1)
	j := k % (n - 1)
	//fmt.Println(i, j)
	if i&1 == 1 {
		return n - j - 1
	} else {
		return j
	}
}
func valueAfterKSeconds(n int, k int) int {
	const mod = 1_000_000_007
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}
	for i := 0; i < k; i++ {
		for j := 1; j < n; j++ {
			arr[j] = (arr[j] + arr[j-1]) % mod
		}
	}
	//for i := 0; i < k; i++ {
	//	v = (v + 1) * n >> 1 % mod
	//	fmt.Println(v)
	//}
	return arr[n-1]
}

//	func maxTotalReward(rewardValues []int) int {
//		n := len(rewardValues)
//		sort.Ints(rewardValues)
//		rbt := redblacktree.NewWithIntComparator()
//		rbt.Put(0, struct{}{})
//		memo := make([]int, 0, n*(n+1)>>1)
//		memoMap := make(map[int]bool)
//		ans := 0
//		for _, v := range rewardValues {
//			pre, _ := rbt.Floor(v - 1) // 前驱节点
//			cur := pre.Key.(int) + v
//			//fmt.Println(pre, cur)
//			ans = max(ans, cur)
//			m := len(memo)
//			if !memoMap[v] {
//				memoMap[v] = true
//				memo = append(memo, v)
//				rbt.Put(v, struct{}{})
//			}
//			for _, val := range memo[:m] {
//				if val >= v {
//					continue
//				}
//				key := val + v
//				if memoMap[key] {
//					continue
//				}
//				memoMap[key] = true
//				memo = append(memo, key)
//				rbt.Put(key, struct{}{})
//			}
//		}
//		//fmt.Println(ans)
//		return ans
//	}

type bitset struct {
	bs []uint
	us int
}

func (b bitset) left(k int) {
	shift, offset := k/b.us, k%b.us
	// 第一段
	l1 := min(k, b.us-offset)
	b.bs[shift] = b.bs[shift] | (b.bs[0] & (1<<l1 - 1) << offset)
	// 中间段
	for i := 0; i < (k-b.us+offset)/b.us; i++ {
		b.bs[shift+i+1] = b.bs[shift+i+1] | (b.bs[i]>>l1 | b.bs[i+1]<<offset)
	}
	// 第三段
	if shift == 0 && offset < b.us>>1 {
		return
	}
	l3 := (offset << 1) % b.us
	if l3 <= offset { // offset 为 bs[shift] 在 k 截断后的前缀
		if shift<<1+1 >= len(b.bs) {
			return
		}
		b.bs[shift<<1+1] = b.bs[shift<<1+1] | b.bs[shift]>>(offset-l3)&(1<<l3-1)
	} else { // 向 shift-1 借 l3-offset 位，然后拼到低位
		b.bs[shift<<1] |= b.bs[shift]&(1<<offset-1)<<(l3-offset) | (b.bs[shift-1] >> (b.us - l3 + offset))
	}
}

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

const w = bits.UintSize

type bitset2 []uint

// b <<= k
func (b bitset2) lsh(k int) bitset2 {
	shift, offset := k/w, k%w
	if offset == 0 {
		// Fast path
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
func (b bitset2) resetRange(start int) bitset2 {
	i := start / w
	//j := start % w
	//fmt.Println(start, j, ^(^uint(0) << j))
	// bug：^uint(0) << 7
	//b[i] &= ^(^uint(0) << (start % w))
	b[i] &= uint(1)<<(start%w) - 1
	clear(b[i+1:])
	return b
}

// b |= c
func (b bitset2) unionFrom(c bitset2) {
	for i, v := range c {
		b[i] |= v
	}
}

func (b bitset2) lastIndex1() int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			return i*w | (bits.Len(b[i]) - 1)
		}
	}
	return -1
}

func maxTotalReward2(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues) // 去重
	m := rewardValues[len(rewardValues)-1]
	f := make(bitset2, (m*2+1)/w+1)
	f[0] = 1
	for _, val := range rewardValues {
		f.unionFrom(slices.Clone(f).lsh(val).resetRange(val * 2))
		fmt.Println(val, f.lastIndex1())
	}
	return f.lastIndex1()
}
