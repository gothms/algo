package _1_basic

/*
既然递归的函数值返回过程和基于循环的迭代法一致，我们直接用迭代法不就好了，为什么还要用递归的数学思想和编程方法呢？
	这是因为，在某些场景下，递归的解法比基于循环的迭代法更容易实现

如何在限定总和的情况下，求所有可能的加和方式？

如何把复杂的问题简单化？
	如何将数学归纳法的思想泛化成更一般的情况？数学归纳法考虑了两种情况：
		1.假设 n=k-1 的时候，问题已经解决（或者已经找到解）。那么只要求解 n=k 的时候，问题如何解决（或者解是多少）
		2.初始状态，就是 n=1 的时候，问题如何解决（或者解 是多少）
	这种思想就是将复杂的问题，每次都解决一点点，并将剩下的任务转化成为更简单的问题等待下次求解，如此反复，直到最简单的形式

递归和循环其实都是迭代法的实现，而且在某些场合下，它们的实现是可以相互转化的。但是，对于某些应用场景，递归确实很难被循环取代。我觉得主要有两点原因：
	第一，递归的核心思想和数学归纳法类似，并更具有广泛性。这两者的类似之处体现在：将当前的问题化解为两部分：一个当前所采取的步骤和另一个更简单的问题
		一个当前所采取的步骤
		另一个更简单的问题
	第二，递归会使用计算机的函数嵌套调用。而函数的调用本身，就可以保存很多中间状态和变量值，因此极大的方便了编程的处理
		正是如此，递归在计算机编程领域中有着广泛的应用，而不仅仅局限在求和等运算操作上

思考
	一个整数可以被分解为多个整数的乘积，例如，6 可以分解为 2x3
	请使用递归编程的方法，为给定的整数 n，找到所有可能的分解（1 在解中最多只能出现 1 次）
	例如，输入 8，输出是可以是 1x8, 8x1, 2x4, 4x2, 1x2x2x2, 1x2x4, ……
*/
