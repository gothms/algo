//某店铺将用于组成套餐的商品记作字符串 goods，其中 goods[i] 表示对应商品。请返回该套餐内所含商品的 全部排列方式 。
//
// 返回结果 无顺序要求，但不能含有重复的元素。
//
//
//
// 示例 1:
//
//
//输入：goods = "agew"
//输出：["aegw","aewg","agew","agwe","aweg","awge","eagw","eawg","egaw","egwa",
//"ewag","ewga","gaew","gawe","geaw","gewa","gwae","gwea","waeg","wage","weag","wega",
//"wgae","wgea"]
//
//
//
//
// 提示：
//
//
// 1 <= goods.length <= 8
//
//
//
//
// Related Topics 字符串 回溯 👍 723 👎 0

package main

import "fmt"

func main() {
	goods := "agew"
	goods = "aab"
	order := goodsOrder(goods)
	fmt.Println(order)
}

// leetcode submit region begin(Prohibit modification and deletion)
func goodsOrder(goods string) []string {
	// 有重复元素
	n := len(goods)
	ans, path := make([]string, 0), []uint8(goods)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		memo := make(map[uint8]struct{})
		for j := i; j < n; j++ {
			if _, ok := memo[path[j]]; !ok {
				memo[path[j]] = struct{}{}
				path[i], path[j] = path[j], path[i]
				dfs(i + 1)
				path[i], path[j] = path[j], path[i]
			}
		}
	}
	dfs(0)
	return ans

	// 无重复元素
	//n := len(goods)
	//ans := make([]string, 0, func(k int) int {
	//	v := 1
	//	for i := 2; i <= k; i++ {
	//		v *= i
	//	}
	//	return v
	//}(n))
	//path := []uint8(goods)
	//var dfs func(int)
	//dfs = func(i int) {
	//	if i == n {
	//		ans = append(ans, string(path))
	//		return
	//	}
	//	dfs(i + 1)
	//	for j := i + 1; j < n; j++ {
	//		path[i], path[j] = path[j], path[i]
	//		dfs(i + 1)
	//		path[i], path[j] = path[j], path[i]
	//	}
	//}
	//dfs(0)
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
