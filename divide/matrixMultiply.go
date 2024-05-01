package divide

import (
	"fmt"
	"math/rand"
)

var aMM, bMM [][]int

var idxMM int

//var cnt int

func matrix() ([][]int, [][]int) { // 模拟点集
	idxMM ^= 1 // arr 用两次
	if idxMM == 0 {
		return aMM, bMM
	}
	//if idx != 0 { // arr 用三次
	//	idx = (idx + 1) % 3
	//	return arrMM
	//}
	//idx++

	//once.Do(func() {
	k := 2
	n := 1 << k
	aMM, bMM = make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		aMM[i], bMM[i] = make([]int, n), make([]int, n)
	}
	seed := 10
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := rand.Intn(seed)
			y := rand.Intn(seed)
			aMM[i][j], bMM[i][j] = x, y
		}
	}
	//})
	return aMM, bMM
}

func printMM(arr [][]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

/*
矩阵乘法
	有两个 n*n 的矩阵 A，B，如何快速求解两个矩阵的乘积 C=A*B？
	https://oi-wiki.org/math/linear-algebra/matrix/
*/

// ====================施特拉森算法====================

// MatrixMultiplyStrassen 施特拉森算法实现矩阵相乘（gpt 实现，有 bug）
// a b 是 2^n * 2^n 矩阵
// wikipedia 施特拉森算法：只有七个矩阵乘法运算，O(n^(log(2,7))) = O(n^2.807)
// 现时时间复杂度最低的矩阵乘法算法是 Coppersmith-Winograd 方法的一种扩展方法，其算法复杂度为 O(n^2.3727)
// 周任飞参与撰写的 2024 年 1 月的论文，降低到了 2.371552：https://zhuanlan.zhihu.com/p/685958403
// 实现：https://github.com/alibaba/MNN/blob/master/source/backend/cpu/compute/StrassenMatmulComputor.cpp
func MatrixMultiplyStrassen(A, B [][]int) [][]int {
	n := len(A)
	if n == 1 {
		C := make([][]int, 1)
		C[0] = make([]int, 1)
		C[0][0] = A[0][0] * B[0][0]
		return C
	}

	// 将矩阵分割成四个子矩阵
	A11, A12, A21, A22 := partitionMatrix(A)
	B11, B12, B21, B22 := partitionMatrix(B)

	// 计算子矩阵的七个乘法结果
	M1 := MatrixMultiplyStrassen(addMatrix(A11, A22), addMatrix(B11, B22))
	M2 := MatrixMultiplyStrassen(addMatrix(A21, A22), B11)
	M3 := MatrixMultiplyStrassen(A11, subMatrix(B12, B22))
	M4 := MatrixMultiplyStrassen(A22, subMatrix(B21, B11))
	M5 := MatrixMultiplyStrassen(addMatrix(A11, A12), B22)
	M6 := MatrixMultiplyStrassen(subMatrix(A21, A11), addMatrix(B11, B12))
	M7 := MatrixMultiplyStrassen(subMatrix(A12, A22), addMatrix(B21, B22))

	// 计算结果矩阵的四个子矩阵
	C11 := addMatrix(subMatrix(addMatrix(M1, M4), M5), M7)
	C12 := addMatrix(M3, M5)
	C21 := addMatrix(M2, M4)
	C22 := addMatrix(subMatrix(addMatrix(M1, M3), M2), M6)

	// 将四个子矩阵组合成结果矩阵
	C := combineMatrix(C11, C12, C21, C22)

	return C
}

// partitionMatrix 将矩阵分割成四个子矩阵
func partitionMatrix(A [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(A)
	n2 := n / 2

	A11 := make([][]int, n2)
	A12 := make([][]int, n2)
	A21 := make([][]int, n2)
	A22 := make([][]int, n2)

	for i := 0; i < n2; i++ {
		A11[i] = make([]int, n2)
		A12[i] = make([]int, n2)
		A21[i] = make([]int, n2)
		A22[i] = make([]int, n2)

		for j := 0; j < n2; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+n2]
			A21[i][j] = A[i+n2][j]
			A22[i][j] = A[i+n2][j+n2]
		}
	}

	return A11, A12, A21, A22
}

// addMatrix 将两个矩阵相加
func addMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)

	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}

	return C
}

// subMatrix 将矩阵B从矩阵A中减去
func subMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)

	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}

	return C
}

// combineMatrix 将四个子矩阵组合成一个矩阵
func combineMatrix(A11, A12, A21, A22 [][]int) [][]int {
	n := len(A11)
	n2 := 2 * n

	C := make([][]int, n2)

	for i := 0; i < n2; i++ {
		C[i] = make([]int, n2)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A11[i][j]
			C[i][j+n] = A12[i][j]
			C[i+n][j] = A21[i][j]
			C[i+n][j+n] = A22[i][j]
		}
	}

	return C
}

// ====================分治算法====================

// MatrixMultiplyDivide 分治算法实现矩阵相乘
func MatrixMultiplyDivide(A, B [][]int) [][]int {
	//A, B = matrix()

	n := len(A)
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}
	strassenHelper(A, B, C, 0, 0, 0, 0, 0, 0, n)
	return C
}

// strassenHelper 使用Strassen算法的辅助函数
func strassenHelper(A, B, C [][]int, aRow, aCol, bRow, bCol, cRow, cCol, size int) {
	if size == 1 {
		C[cRow][cCol] += A[aRow][aCol] * B[bRow][bCol]
		return
	}

	newSize := size / 2

	// 子矩阵相加
	strassenHelper(A, B, C, aRow, aCol, bRow, bCol, cRow, cCol, newSize)
	strassenHelper(A, B, C, aRow, aCol+newSize, bRow+newSize, bCol, cRow, cCol, newSize)

	// 子矩阵相乘
	strassenHelper(A, B, C, aRow, aCol, bRow, bCol+newSize, cRow, cCol+newSize, newSize)
	strassenHelper(A, B, C, aRow, aCol+newSize, bRow+newSize, bCol+newSize, cRow, cCol+newSize, newSize)

	strassenHelper(A, B, C, aRow+newSize, aCol, bRow, bCol, cRow+newSize, cCol, newSize)
	strassenHelper(A, B, C, aRow+newSize, aCol+newSize, bRow+newSize, bCol, cRow+newSize, cCol, newSize)

	strassenHelper(A, B, C, aRow+newSize, aCol, bRow, bCol+newSize, cRow+newSize, cCol+newSize, newSize)
	strassenHelper(A, B, C, aRow+newSize, aCol+newSize, bRow+newSize, bCol+newSize, cRow+newSize, cCol+newSize, newSize)
}

// ====================暴力====================

// MatrixMultiplyForce 矩阵相乘，暴力法
// 对于比较小的矩阵，可以考虑直接手动展开循环以减小常数
// O(n^(log(2,8))) = O(n^3)
func MatrixMultiplyForce(a, b [][]int) [][]int {
	//a, b = matrix()

	n := len(a) // n*n
	c := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	c[i] = make([]int, n)
	//	for j := 0; j < n; j++ {
	//		for k := 0; k < n; k++ {
	//			c[i][j] += a[i][k] * b[k][j]
	//		}
	//	}
	//}
	// 优化：重新排列循环以提高空间局部性，不会改变矩阵乘法的时间复杂度，但是会在得到常数级别的提升
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for k := 0; k < n; k++ {
			v := a[i][k]
			for j := 0; j < n; j++ {
				c[i][j] += v * b[k][j]
			}
		}
	}
	return c
}
