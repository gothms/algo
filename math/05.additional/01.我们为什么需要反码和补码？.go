package _5_additional

/*
什么是符号位？为什么要有符号位？
	用一句话来概括就是，符号位是有符号二进制数中的最高位，我们需要它来表示负数

	在实际的硬件系统中，计算机 CPU 的运算器只实现了加法器，而没有实现减法器
		那么计算机如何做减法呢？我们可以通过加上一个负数来达到这个目的
	如何让计算机理解哪些是正数，哪些是负数呢？为此，人们把二进制数分为有符号数（signed）和无符号数（unsigned）
什么是溢出？
	对于 n 位的数字类型，符号位是 1，后面 n-1 位全是 0，我们把这种情形表示为 -2^(n-1) ，而不是 2^(n-1)
		一旦某个数字超过了这些限定，就会发生溢出。如果超出上限，就叫上溢出（overflow）。如果超出了下限，就叫下溢出（underflow）
	那么溢出之后会发生什么呢？
		上溢出之后，又从下限开始，最大的数值加 1，就变成了最小的数值，周而复始，这不就是余数和取模的概念吗？
		所以说，计算机数据的溢出，就相当于取模
		而用于取模的除数就是数据类型的上限减去下限的值，再加上 1，也就是 (2^(n-1)-1)-(-2^(n-1))+1=2x2^(n-1)-1+1=2^n-1+1
		这个除数为什么不直接写成 2^n 呢？这是因为 2^n 已经是 n+1 位了，已经超出了 n 位所能表示的范围
二进制的原码、反码及补码
	原码就是我们看到的二进制的原始表示
		对于有符号的二进制来说，原码的最高位是符号位，而其余的位用来表示该数字绝对值的二进制
		所以 +2 的原码是 000…010，-2 的的原码是 100…010

		那么我们是不是可以直接使用负数的原码来进行减法计算呢？答案是否定的。我还是以 3+(-2) 为例
			相加后的结果是二进制 100…0101，它的最高位是 1，表示负数，而最低的 3 位是 101，表示 5，所以结果就是 -5 的原码了，而 3+(-2) 应该等于 1，两者不符
	如果负数的原码并不适用于减法操作，那该怎么办呢？这个问题的解答还要依赖计算机的溢出机制
		i-j=(i-j)+(2^n-1+1)=i+(2^n-1-j+1)
		有了反码的定义，那么就可以得出 i-j=i+(2^n-1-j+1)=i 的原码 +(-j 的反码)+1
		如果我们把 -j 的反码加上 1 定义为 -j 的补码，就可以得到 i-j=i 的原码 +(-j 的补码)
		由于正数的加法无需负数的加法这样的变换，因此正数的原码、反码和补码三者都是一样的。最终，我们可以得到 i-j=i 的补码 +(-j 的补码)
	换句话说，计算机可以通过补码，正确地运算二进制减法
		再来用 3+(-2) 来验证一下。正数 3 的补码仍然是 0000…0011，-2 的补码是 1111…1110，两者相加，最后得到了正确的结果 1 的二进制
		可见，溢出本来是计算机数据类型的一种局限性，但在负数的加法上，它倒是可以帮我们大忙
*/
