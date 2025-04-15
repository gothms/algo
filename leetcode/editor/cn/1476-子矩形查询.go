package main

import "slices"

func main() {

}

/*
思路一：二分
	使用二分查找，新的区间可能将老区间分割
	如果是二维数组，那么分割的情况需要分几种情况讨论，且在二维数组上二分查找区间
思路二：线段树
	假如是一维的数组，可以使用线段树懒更新区间
	那么二维的数组，使用线段树如何解决呢？
*/

// leetcode submit region begin(Prohibit modification and deletion)
type SubrectangleQueries struct {
	g      [][]int
	update [][5]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle, make([][5]int, 0)}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.update = append(this.update, [5]int{row1, col1, row2, col2, newValue})
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for _, u := range slices.Backward(this.update) {
		if u[0] <= row && u[2] >= row && u[1] <= col && u[3] >= col {
			return u[4]
		}
	}
	return this.g[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
//leetcode submit region end(Prohibit modification and deletion)
