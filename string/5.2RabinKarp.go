package string

/*
Rabin-Karp 字符串匹配算法
	hash 匹配
	平均O(n)

代码注释：
	p：模式串
	s：文本字符串
	1.计算模式串 p 的hash值，和 primeRK^len(p)
		1.1.计算模式串 p 的hash值
		1.2.计算 primeRK^len(p)
	2.根据Rabin-Karp算法，匹配字符串
		2.1.计算s前n的hash值，n为p的长度
		2.2.第一次hash匹配
		2.3.匹配所有hash，从第n个字符开始（因为0~n-1已匹配）
		2.4.匹配成功
			二次检查：s[i-n:i] =? p
	补充：2.3
		Rabin-Karp的hash算法
		例abcd：假设长度为3
		hash_abc = a * primeRK^2 + a * primeRK^1 + a * primeRK^0
		hash_bcd = hash_abc * primeRK + d * primeRK^0 - a * primeRK^3
	补充：1.2
		计算了 pow^i，相比 i-- 循环减少了次数
		使算法更加的简便高效
*/
//const primeRK = 16777619

// RabinKarp 字符串匹配算法
func RabinKarp(s, p string) []int {
	ans, hash, n := make([]int, 0), uint32(0), len(p)
	hashP, pow := rkHelper(p) // 1
	for i := 0; i < n; i++ {  // 2
		hash = hash*primeRK + uint32(s[i]) // 2.1
	}
	if hash == hashP && s[:n] == p { // 2.2
		ans = append(ans, 0)
	}
	for i := n; i < len(s); { // 2.3
		hash = hash*primeRK + uint32(s[i]) - pow*uint32(s[i-n])
		i++
		if hash == hashP && s[i-n:i] == p { // 2.4
			ans = append(ans, i-n)
		}
	}
	return ans
}
func rkHelper(p string) (uint32, uint32) { // 1
	var hash, pow, base uint32 = 0, 1, primeRK
	for i := 0; i < len(p); i++ { // 1.1
		hash = hash*base + uint32(p[i])
	}
	for i := len(p); i > 0; i >>= 1 { // 1.2
		if i&1 != 0 {
			pow *= base
		}
		base *= base
	}
	return hash, pow
}
