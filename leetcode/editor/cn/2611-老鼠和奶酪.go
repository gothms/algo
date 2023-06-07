//有两只老鼠和 n 块不同类型的奶酪，每块奶酪都只能被其中一只老鼠吃掉。
//
// 下标为 i 处的奶酪被吃掉的得分为：
//
//
// 如果第一只老鼠吃掉，则得分为 reward1[i] 。
// 如果第二只老鼠吃掉，则得分为 reward2[i] 。
//
//
// 给你一个正整数数组 reward1 ，一个正整数数组 reward2 ，和一个非负整数 k 。
//
// 请你返回第一只老鼠恰好吃掉 k 块奶酪的情况下，最大 得分为多少。
//
//
//
// 示例 1：
//
//
//输入：reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2
//输出：15
//解释：这个例子中，第一只老鼠吃掉第 2 和 3 块奶酪（下标从 0 开始），第二只老鼠吃掉第 0 和 1 块奶酪。
//总得分为 4 + 4 + 3 + 4 = 15 。
//15 是最高得分。
//
//
// 示例 2：
//
//
//输入：reward1 = [1,1], reward2 = [1,1], k = 2
//输出：2
//解释：这个例子中，第一只老鼠吃掉第 0 和 1 块奶酪（下标从 0 开始），第二只老鼠不吃任何奶酪。
//总得分为 1 + 1 = 2 。
//2 是最高得分。
//
//
//
//
// 提示：
//
//
// 1 <= n == reward1.length == reward2.length <= 10⁵
// 1 <= reward1[i], reward2[i] <= 1000
// 0 <= k <= n
//
//
// Related Topics 贪心 数组 排序 堆（优先队列） 👍 10 👎 0

package main

import "fmt"

func main() {
	a1 := []int{1, 1, 3, 4}
	a2 := []int{4, 4, 1, 1}
	a1 = []int{1, 1}
	a2 = []int{1, 1}
	a1 = []int{1, 4, 4, 6, 4}
	a2 = []int{6, 5, 3, 6, 1}
	cheese := miceAndCheese(a1, a2, 2)
	fmt.Println(cheese)
}

/*
思路：topK
	1.老鼠1 选择吃掉 reward1 中最优的 k 个蛋糕，怎么选是最优的蛋糕呢？
		reward1[i] - reward2[i] 差值越大，则越优
	2.有了最优选择，最大得分 sum 的策略就明显了：
		2.1.把 reward1 和 reward2 中，同索引的值作一一映射关系
			reward1[i] - reward2[i] 差值是前 k 大，则让 老鼠1 选择，sum+=reward1[i]
			reward1[i] - reward2[i] 差值不是前 k 大，则让 老鼠2 选择，sum+=reward2[i]
		2.2.可以将 reward1[i] - reward2[i] 存入一个 cache 数组，i=[0,n)，分两步求 sum
			a)sum 等于 reward2 所有元素之和
			b)从 cache 中获取 topK 个元素，sum 再加上这 topK 个元素
	3.使用 quickSort 求 topK
*/
//leetcode submit region begin(Prohibit modification and deletion)
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	// topK
	sum, n := 0, len(reward1)
	if k == 0 { // 全是 mice_2 吃蛋糕
		for i := 0; i < n; i++ {
			sum += reward2[i]
		}
		return sum
	}
	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i] = reward1[i] - reward2[i]
		sum += reward2[i] // a)reward2 元素和
	}
	kQuickSort(cache, 0, n-1, k)
	for i := 0; i < k; i++ {
		sum += cache[i] // b)加上 tokK
	}
	return sum
}
func kQuickSort(arr []int, l, r, k int) {
	if l >= r {
		return
	}
	pivot := kPartition(arr, l, r)
	switch {
	case pivot < k:
		kQuickSort(arr, pivot+1, r, k)
	case pivot > k:
		kQuickSort(arr, l, pivot-1, k)
	}
}
func kPartition(arr []int, l, r int) int {
	pivot, counter := l, l+1
	for i := l + 1; i <= r; i++ {
		if arr[i] >= arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	counter--
	arr[counter], arr[pivot] = arr[pivot], arr[counter]
	return counter
}

//leetcode submit region end(Prohibit modification and deletion)
