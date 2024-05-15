package main

import (
	"algo/oiWiki"
	"fmt"
)

func main() {
	// 平面最近点对
	//distance := divide.NearestPointsDivide(nil)
	////fmt.Printf("NearestPointsDivide: %.4f\n", distance)
	//d := divide.NearestPointsNotDivide(nil)
	////fmt.Printf("NearestPointsDivide: %.4f\n", d)
	////force := divide.NearestPointsForce(nil)
	//if distance != d {
	//	//fmt.Println("error: ", distance, d, force)
	//	fmt.Println("error: ", distance, d)
	//}
	//
	//for i := 0; i < 10000; i++ {
	//	x := divide.NearestPointsDivide(nil)
	//	y := divide.NearestPointsNotDivide(nil)
	//	//z := divide.NearestPointsForce(nil)
	//	if x != y {
	//		//fmt.Println("error: ", x, y, z)
	//		fmt.Println("error: ", x, y)
	//	}
	//}
	//divide.OverNP()

	// 矩阵乘法
	A := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	B := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	mmf := oiWiki.MatrixMultiplyForce(A, B)
	fmt.Println(mmf)
	mms := oiWiki.MatrixMultiplyDivide(A, B)
	fmt.Println(mms)

	//mms := divide.MatrixMultiplyStrassen(A, B)
	//for _, row := range mms {
	//	fmt.Println(row)
	//}

	for i, m := range mmf {
		for j, v := range m {
			if mms[i][j] != v {
				fmt.Println(i, j, v, mms[i][j])
			}
		}
	}
}
