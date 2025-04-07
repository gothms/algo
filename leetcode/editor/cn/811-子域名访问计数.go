package main

import "strconv"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func subdomainVisits(cpdomains []string) []string {
	memo := make(map[string]int)
	for _, cp := range cpdomains {
		var cnt int
		for i, c := range cp {
			switch c {
			case ' ':
				cnt, _ = strconv.Atoi(cp[:i])
				//memo[cp[i+1:]] += cnt
				fallthrough
			case '.':
				memo[cp[i+1:]] += cnt
			}
		}
	}
	ans := make([]string, 0, len(memo))
	for k, v := range memo {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
