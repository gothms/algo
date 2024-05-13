//给你一个数组 nums ，请你完成两类查询。
//
//
// 其中一类查询要求 更新 数组 nums 下标对应的值
// 另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
//
//
// 实现 NumArray 类：
//
//
// NumArray(int[] nums) 用整数数组 nums 初始化对象
// void update(int index, int val) 将 nums[index] 的值 更新 为 val
// int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元
//素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
//
//
//
//
// 示例 1：
//
//
//输入：
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//输出：
//[null, 9, null, 8]
//
//解释：
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
//numArray.update(1, 2);   // nums = [1,2,5]
//numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -100 <= nums[i] <= 100
// 0 <= index < nums.length
// -100 <= val <= 100
// 0 <= left <= right < nums.length
// 调用 update 和 sumRange 方法次数不大于 3 * 10⁴
//
//
// Related Topics 设计 树状数组 线段树 数组 👍 695 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	arr []int
	bit []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	bit := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bit[i] += nums[i-1]
		if j := i + i&-i; j <= n {
			bit[j] += bit[i]
		}
	}
	return NumArray{nums, bit}
}
func (this *NumArray) Update(index int, val int) {
	delta := val - this.arr[index]
	this.arr[index] = val
	index++
	for ; index < len(this.bit); index += index & -index {
		this.bit[index] += delta
	}
}
func (this *NumArray) prefixSum(i int) int {
	var sum int
	for ; i > 0; i &= i - 1 {
		sum += this.bit[i]
	}
	return sum
}
func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum(right+1) - this.prefixSum(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

//type NumArray struct {
//	arr []int
//	bit []int
//}
//
//func Constructor(nums []int) NumArray {
//	return NumArray{nums, func() []int {
//		n := len(nums)
//		b := make([]int, n)
//		for i, j := 0, 0; i < n; i++ {
//			b[i] += nums[i]
//			if j = i | (i + 1); j < n {
//				b[j] += b[i]
//			}
//		}
//		return b
//	}()}
//}
//func (this *NumArray) Update(index int, val int) {
//	n, d := len(this.bit), val-this.arr[index]
//	this.arr[index] = val
//	for i := index; i < n; i |= i + 1 {
//		this.bit[i] += d
//	}
//}
//func (this *NumArray) SumRange(left int, right int) int {
//	return this.prefixSum(right) - this.prefixSum(left-1)
//}
//func (this *NumArray) prefixSum(i int) int {
//	sum := 0
//	for ; i >= 0; i -= ^i & (i + 1) {
//		sum += this.bit[i]
//	}
//	return sum
//}
