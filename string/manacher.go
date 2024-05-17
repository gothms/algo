package str

/*
Manacher 算法：https://oi-wiki.org/string/manacher/
寻找所有奇数长度子回文串的情况，即只计算 d1[]，寻找所有偶数长度子回文串的算法（即计算数组 d2[]）
	维护 l r
		维护已找到的最靠右的子回文串的边界 (l, r)（即具有最大 r 值的回文串，其中 l 和 r 分别为该回文串左右边界的位置）
		置 l = 0 和 r = -1（-1需区别于倒序索引位置，这里可为任意负数，仅为了循环初始时方便）
	示例
		s = "bananas" d1 = [1 1 2 3 2 1 1]
		s = "dabbac" d2 = [0 0 0 2 0 0]
	核心逻辑
		k = min(d1[l+r-i], r-i+1)
	1.考察值 d1[j]：d1[l+r-i]
		在子回文串 (l, r) 中反转位置 i，即得到 j = l + (r - i)，现在来考察值 d1[j]
		因为位置 j 同位置 i 对称，几乎总是可以置 d1[i] = d1[j]
		可认为以 j 为中心的回文串被「拷贝」至以 i 为中心的位置上
		...s[l]...s[j-d1[j]+1]...s[j]...s[j+d1[j]-1]...s[i-d1[j]+1]...s[i]...s[i+d1[j]-1]...s[r]...
	2.特殊情况：r-i+1
		当「内部」的回文串到达「外部」回文串的边界时，即 j - d1[j] + 1 <= l（或者等价的说，i + d1[j] - 1 >= r）
		因为在「外部」回文串范围以外的对称性没有保证，因此直接置 d1[i] = d1[j] 将是不正确的：我们没有足够的信息来断言在位置 i 的回文串具有同样的长度
		为了正确处理这种情况，应该「截断」回文串的长度，即置 d1[i] = r - i。之后将运行朴素算法以尝试尽可能增加 d1[i] 的值
统一处理
	本质上即求 s' 的 d1[]，因此在得到 s' 后，代码同计算 d1[] 的一样
		给定一个长度为 n 的字符串 s，在其 n + 1 个空中插入分隔符 #，从而构造一个长度为 2n + 1 的字符串 s'
		对于字母间的 #，其实际意义为 s 中对应的「空」。而两端的 # 则是为了实现的方便
	易得，在 s' 中，d1[i] 表示在 s 中以对应位置为中心的极大子回文串的 总长度加一
		设 m 为偶数
		奇数长度回文串：对于 s 中一个以字母为中心的极大子回文串，设其长度为 m + 1，则其在 s' 中对应一个以相应字母为中心，长度为 2m + 3 的极大子回文串
		偶数长度回文串：而对于 s 中一个以空为中心的极大子回文串，设其长度为 m，则其在 s' 中对应一个以相应表示空的 # 为中心，长度为 2m + 1 的极大子回文串

lc
	5
*/

// manacher 统一处理奇数、偶数长度子回文串
func manacher(s string) string {
	n := len(s)<<1 + 1
	str, mnc := make([]int32, 0, n), make([]int, n)
	str = append(str, '#')
	for _, c := range s {
		str = append(str, c, '#')
	}
	l, r := 0, -1
	for i := range str {
		var k int
		if i > r {
			k = 1
		} else { // i <= r
			k = min(mnc[l+r-i], r-i+1)
		}
		for i-k >= 0 && i+k < n && str[i-k] == str[i+k] {
			k++
		}
		mnc[i] = k
		k--
		if i+k > r {
			l = i - k
			r = i + k
		}
	}
	idx, cnt := 0, 0
	for i, c := range mnc {
		if c > cnt {
			idx, cnt = i, c // idx 为偶数，则为 '#'
		}
	}
	l = (idx - cnt + 1) >> 1 // cnt - 1：回文子串长度，(idx - cnt + 1) >> 1：在原字符串 s 中的起始位置
	return s[l : l+cnt-1]
}

// manacher_ 分别寻找奇数、偶数长度子回文串
func manacher_(s string) string {
	n := len(s)
	d1 := make([]int, n)
	for i, l, r := 0, 0, -1; i < n; i++ {
		k := 0
		if i > r {
			k = 1
		} else { // i <= r：尝试从已计算过的 d1[] 的值中获取一些信息
			k = min(d1[l+r-i], r-i+1)
		}
		for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
			k++ // 朴素算法
		}
		d1[i] = k // 半径长度
		k--
		if i+k > r { // 更新 l r
			l = i - k
			r = i + k // 已找到的最靠右的子回文串的右边界
		}
	}
	d2 := make([]int, n) // 计算偶数长度回文串数组 d2[] 的算法同上述计算奇数长度回文串数组 d1[] 的算法十分类似
	for i, l, r := 0, 0, -1; i < n; i++ {
		k := 0
		if i > r {
			k = 0
		} else {
			k = min(d2[l+r-i+1], r-i+1)
		}
		for i-k-1 >= 0 && i+k < n && s[i-k-1] == s[i+k] {
			k++
		}
		d2[i] = k
		k--
		if i+k > r {
			l = i - k - 1
			r = i + k
		}
	}

	maxL, idx := 1, 1 // 返回最长回文子串（最后一个）
	for i := 0; i < n; i++ {
		if v1, v2 := d1[i]<<1-1, d2[i]<<1; v1 > maxL {
			maxL, idx = v1, i+d1[i]
		} else if v2 > maxL {
			maxL, idx = v2, i+d2[i]
		}
	}
	return s[idx-maxL : idx]
}
