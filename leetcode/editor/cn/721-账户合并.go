package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func accountsMerge(accounts [][]string) [][]string {
	ans := make([][]string, 0)
	emailToIdx := map[string][]int{}
	for i, account := range accounts {
		for _, email := range account[1:] {
			emailToIdx[email] = append(emailToIdx[email], i)
		}
	}

	vis := make([]bool, len(accounts))
	emailSet := map[string]struct{}{} // 用于收集 DFS 中访问到的邮箱地址
	var dfs func(int)
	dfs = func(i int) {
		vis[i] = true
		for _, email := range accounts[i][1:] {
			if _, has := emailSet[email]; has {
				continue
			}
			emailSet[email] = struct{}{}
			for _, j := range emailToIdx[email] { // 遍历所有包含该邮箱地址的账户下标 j
				if !vis[j] { // j 没有访问过
					dfs(j)
				}
			}
		}
	}

	for i, b := range vis {
		if b {
			continue
		}
		clear(emailSet)
		dfs(i)

		res := make([]string, 1, len(emailSet)+1)
		res[0] = accounts[i][0]
		for email := range emailSet {
			res = append(res, email)
		}
		slices.Sort(res[1:])

		ans = append(ans, res)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
