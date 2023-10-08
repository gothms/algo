//这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i -
//arr[i]。
//
// 请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
//
// 注意，不管是什么情况下，你都无法跳到数组之外。
//
//
//
// 示例 1：
//
// 输入：arr = [4,2,3,0,3,1,2], start = 5
//输出：true
//解释：
//到达值为 0 的下标 3 有以下可能方案：
//下标 5 -> 下标 4 -> 下标 1 -> 下标 3
//下标 5 -> 下标 6 -> 下标 4 -> 下标 1 -> 下标 3
//
//
// 示例 2：
//
// 输入：arr = [4,2,3,0,3,1,2], start = 0
//输出：true
//解释：
//到达值为 0 的下标 3 有以下可能方案：
//下标 0 -> 下标 4 -> 下标 1 -> 下标 3
//
//
// 示例 3：
//
// 输入：arr = [3,0,2,1,2], start = 2
//输出：false
//解释：无法到达值为 0 的下标 1 处。
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 5 * 10^4
// 0 <= arr[i] < arr.length
// 0 <= start < arr.length
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 👍 156 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func canReach(arr []int, start int) bool {
	// bfs
	n := len(arr)
	visited := make([]bool, n)
	visited[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]
		if r := f + arr[f]; r < n && !visited[r] { // 往右
			if arr[r] == 0 {
				return true
			}
			visited[r] = true
			queue = append(queue, r)
		}
		if l := f - arr[f]; l >= 0 && !visited[l] { // 往左
			if arr[l] == 0 {
				return true
			}
			visited[l] = true
			queue = append(queue, l)
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
