//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：可以设计并实现时间复杂度为 O(n) 的解决方案吗？
//
//
//
//
// 注意：本题与主站 128 题相同： https://leetcode-cn.com/problems/longest-consecutive-
//sequence/
//
// Related Topics 并查集 数组 哈希表 👍 83 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	// 个人：并查集
	//type uni struct {
	//	p, cap int
	//}
	//memo := make(map[int]*uni)
	//var find func(int) *uni
	//find = func(p int) *uni {
	//	if v := memo[p].p; v != p {
	//		memo[p] = find(v)
	//	}
	//	return memo[p]
	//}
	//union := func(p, q int) *uni {
	//	rootP, rootQ := find(p), find(q)
	//	if rootP != rootQ { // 比较指针
	//		rootP.cap += rootQ.cap // 合并区间
	//		memo[q] = rootP
	//	}
	//	return rootP
	//}
	//ans := 0
	//for _, v := range nums {
	//	if memo[v] != nil {
	//		continue
	//	}
	//	l, r := memo[v-1], memo[v+1]
	//	var u *uni
	//	if l != nil && r != nil {
	//		u = union(l.p, r.p)
	//	} else if l != nil {
	//		u = find(l.p)
	//	} else if r != nil {
	//		u = find(r.p)
	//	} else {
	//		u = &uni{v, 0}
	//	}
	//	u.cap++
	//	memo[v] = u
	//	ans = max(ans, u.cap)
	//}
	//return ans

	// lc：hash
	memo := make(map[int]struct{})
	for _, v := range nums {
		memo[v] = struct{}{}
	}
	var ans int
	for _, v := range nums {
		if _, ok := memo[v-1]; ok { // 有 v-1 存在
			continue
		}
		cnt, cur := 1, v+1
		for _, ok := memo[cur]; ok; _, ok = memo[cur] { // 往 v+1 找
			cnt++
			cur++
		}
		ans = max(ans, cnt)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
