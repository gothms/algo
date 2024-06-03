package main

import (
	"fmt"
)

func main() {
	s := "babab" // 9
	//s = "azbazbzaz" // 14
	//s = "xklmswdwwpyavhsbmkpwknawisfihbnoibjepjzpgklopftcryrsrhhybdj"                    // 59
	//s = "kjjopkrzvvihbramxqnolmhqzutepxouqsnoofpqnuceiomqkriwcxzriiziixrkiskewchkq"      // 78
	//s = "fmfpsvagoyzarkzolszfxuglyygaikqasgcjxnkuenlirdqcpftuxginsrybdgqyiuqwlipnsowfub" // 82
	scores := sumScores(s)
	fmt.Println(scores)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sumScores(s string) int64 {
	// 扩展 KMP
	ans, n := 0, len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		//if i <= r && z[i-l] < r-i+1 {
		//	fmt.Println(i, l, r)
		//	z[i] = z[i-l]
		//} else {
		//	z[i] = max(0, r-i+1)
		//	for i+z[i] < n && s[z[i]] == s[i+z[i]] {
		//		z[i]++
		//	}
		//}
		//if i+z[i]-1 > r {
		//	l, r = i, i+z[i]-1
		//}
		//ans += z[i]

		z[i] = max(0, min(z[i-l], r-i+1))
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		ans += z[i]
	}
	return int64(ans + n)

	// 后缀数组
	//n := len(s)
	//sa := *(*[]int32)(unsafe.Pointer(
	//	reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	//rk, height := make([]int, n), make([]int, n)
	//for i, v := range sa {
	//	rk[v] = i
	//}
	//for i, k := 0, 0; i < n; i++ {
	//	if rk[i] == 0 {
	//		continue
	//	}
	//	if k > 0 {
	//		k--
	//	}
	//	for j := int(sa[rk[i]-1]); i+k < n && j+k < n && s[i+k] == s[j+k]; {
	//		k++
	//	}
	//	height[rk[i]] = k
	//}
	//var ans int
	//for i, l := rk[0], n; i >= 0 && height[i] > 0; i-- { // 更低排名
	//	l = min(l, height[i]) // 可能超过 height[i]，也防止 s[int(sa[i-1])+l-1] 越界
	//	for s[l-1] != s[int(sa[i-1])+l-1] {
	//		l-- // 公共前缀只会越来越短
	//	}
	//	ans += l
	//}
	//for i, r := rk[0]+1, n; i < n && height[i] > 0; i++ { // 更高排名
	//	r = min(r, height[i])
	//	for s[r-1] != s[int(sa[i])+r-1] {
	//		r--
	//	}
	//	ans += r
	//}
	//return int64(n + ans)
}

//leetcode submit region end(Prohibit modification and deletion)
