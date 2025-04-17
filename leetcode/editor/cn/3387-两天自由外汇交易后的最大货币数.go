package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
	// lc：建图 & dfs
	// 输入保证两个转换图在各自的天数中没有矛盾或循环
	type pair struct {
		t string
		r float64
	}
	calc := func(init string, pairs [][]string, rates []float64) map[string]float64 {
		g := make(map[string][]pair)
		for i, p := range pairs { // 建图
			f, t := p[0], p[1]
			g[f] = append(g[f], pair{t, rates[i]})
			g[t] = append(g[t], pair{f, 1 / rates[i]})
		}
		amount := make(map[string]float64)
		var dfs func(string, float64)
		dfs = func(f string, v float64) {
			amount[f] = v
			for _, p := range g[f] {
				t, r := p.t, p.r
				if amount[t] == 0 {
					dfs(t, v*r)
				}
			}
		}
		dfs(init, 1.0)
		return amount // 从 init 开始，能兑换的最大钱数
	}
	day1 := calc(initialCurrency, pairs1, rates1)
	day2 := calc(initialCurrency, pairs2, rates2)
	var ans float64
	for k, v := range day2 { // day2 作为除数，防止除数 =0
		ans = max(ans, day1[k]/v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
