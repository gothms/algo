//实现支持下列接口的「快照数组」- SnapshotArray：
//
//
// SnapshotArray(int length) - 初始化一个与指定长度相等的 类数组 的数据结构。初始时，每个元素都等于 0。
// void set(index, val) - 会将指定索引 index 处的元素设置为 val。
// int snap() - 获取该数组的快照，并返回快照的编号 snap_id（快照号是调用 snap() 的总次数减去 1）。
// int get(index, snap_id) - 根据指定的 snap_id 选择快照，并返回该快照指定索引 index 的值。
//
//
//
//
// 示例：
//
// 输入：["SnapshotArray","set","snap","set","get"]
//     [[3],[0,5],[],[0,6],[0,0]]
//输出：[null,null,0,null,5]
//解释：
//SnapshotArray snapshotArr = new SnapshotArray(3); // 初始化一个长度为 3 的快照数组
//snapshotArr.set(0,5);  // 令 array[0] = 5
//snapshotArr.snap();  // 获取快照，返回 snap_id = 0
//snapshotArr.set(0,6);
//snapshotArr.get(0,0);  // 获取 snap_id = 0 的快照中 array[0] 的值，返回 5
//
//
//
// 提示：
//
//
// 1 <= length <= 50000
// 题目最多进行50000 次set，snap，和 get的调用 。
// 0 <= index < length
// 0 <= snap_id < 我们调用 snap() 的总次数
// 0 <= val <= 10^9
//
//
// Related Topics 设计 数组 哈希表 二分查找 👍 130 👎 0

package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type SnapshotArray struct {
	snapArr [][][2]int // [2]：val snap_id
	k       int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{make([][][2]int, length), 0}

	// lc
	// Set：不检查，直接添加
	// Get：查 arr[i][1] > snap_id，idx == 0 则返回 0
}

func (this *SnapshotArray) Set(index int, val int) {
	arr := this.snapArr[index]
	if n := len(arr); n > 0 && arr[n-1][1] == this.k { // 同 snap_id，则更新 val
		arr[n-1][0] = val
	} else {
		this.snapArr[index] = append(arr, [2]int{val, this.k}) // 添加 val
	}
}

func (this *SnapshotArray) Snap() int {
	ret := this.k
	this.k++
	return ret
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	arr := this.snapArr[index]
	if len(arr) == 0 {
		return 0
	}
	idx := sort.Search(len(arr), func(i int) bool { // 二分查找 snap_id
		return arr[i][1] > snap_id
	}) - 1
	if idx < 0 { // arr[0][1] > snap_id，此时 idx = -1
		return 0
	}
	return arr[idx][0]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
//leetcode submit region end(Prohibit modification and deletion)
