package maths

/*
https://oi-wiki.org/math/number-theory/inverse/

费马小定理
	若 p 为素数，gcd(a, p) = 1，则 a^(p - 1) ≡ 1 (mod p)
	另一个形式：对于任意整数 a，有 a^p ≡ a (mod p)
欧拉定理

扩展欧拉定理

乘法逆元 / 模逆元
	如果一个线性同余方程 a*x ≡ 1 mod b，则 x 称为 a mod b 的逆元，记作 a^-1。
		a*x % b == 1
	如何求逆元
		扩展欧几里得法
		快速幂法
			快速幂法使用了费马小定理，要求 b 是一个素数；而扩展欧几里得法只要求
		线性求逆元
		线性求任意 n 个数的逆元

示例
	a/b % p ≡ a*(b的逆元) % p
	如 a=6 b=2 p=5，6/2 % 5 ≡ 6*3 % 5
*/

// ====================扩展欧几里得法====================

// ExGcd 扩展欧几里得法，写法一
// x,y 可以是任意初始值，返回值 x 即为逆元，扩展欧几里得法只要求 gcd(a,b)=1
func ExGcd(a, b int, x, y *int) {
	if b == 0 {
		*x, *y = 1, 0
		return
	}
	ExGcd(b, a%b, y, x)
	*y -= a / b * *x
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Inv 扩展欧几里得法，写法二
// mod 即为 b
func Inv(a, mod int) int {
	x, _ := extendGcd(a, mod)
	return (x + mod) % mod
}

func extendGcd(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	x, y := extendGcd(b, a%b)
	return y, x - (a/b)*y
}

// ====================快速幂法====================

// QPow 快速幂法，使用了费马小定理，要求 b 是一个素数；而扩展欧几里得法只要求 gcd(a,b)=1
// x ≡ a^(b-2) (mod b)
func QPow(a, b int) int {
	ans := 1
	a = (a%b + b) % b
	for i := b - 2; i != 0; i >>= 1 {
		if i&1 != 0 {
			ans = a * ans % b
		}
		a = a * a % b
	}
	return ans
}

// ====================线性求逆元====================

// ExGcdN 求出 1,2,...,n 中每个数关于 p 的逆元
func ExGcdN(n, p int) []int {
	inv := make([]int, n+1)
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (p - p/i) * inv[p%i] % p
	}
	return inv[1:]
}

// ====================线性求任意 n 个数的逆元====================

// ExGcdS 求任意给定 n 个数（1 <= a[i] < p）的逆元
// O(n+log p)
func ExGcdS(a []int, p int) []int {
	n := len(a)
	s := make([]int, n+1)
	s[0] = 1
	for i := 1; i <= n; i++ { // 首先计算 n 个数的前缀积，记为 s[i]
		s[i] = s[i-1] * a[i-1] % p
	}
	inv, sv := make([]int, n+1), make([]int, n+1)
	sv[n] = QPow(s[n], p)     // 使用快速幂或扩展欧几里得法计算 s[n] 的逆元，记为 sv[n]。若 qPow 的入参为 b-2，则 sv[n] = qPow(s[n], p-2)
	for i := n; i >= 1; i-- { // 因为 sv[n] 是 n 个数的积的逆元，所以当我们把它乘上 a[n] 时，就会和 a[n] 的逆元抵消，于是就得到了 a[1] 到 a[n-1] 的积逆元，记为 sv[n-1]
		sv[i-1] = sv[i] * a[i-1] % p
	}
	for i := 1; i <= n; i++ { // a[i] 的逆元就可以用 s[i-1]*sv[i] 求得
		inv[i] = sv[i] * s[i-1] % p
	}
	return inv[1:]
}
