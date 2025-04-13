package main

import (
	"fmt"
	"slices"
)

func main() {
	chars := []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}                              // ["a","3","b","2","a","2"]
	chars = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'} // ["a","b","1","2"]
	i := compress(chars)
	fmt.Printf("%c\n", chars[:i])
}

// leetcode submit region begin(Prohibit modification and deletion)
func compress(chars []byte) int {
	// 个人
	getCnt := func(c int) []byte {
		if c == 1 {
			return nil
		}
		var ret []byte
		for ; c > 0; c /= 10 {
			ret = append(ret, byte(c%10+'0'))
		}
		slices.Reverse(ret)
		return ret
	}
	ans, cnt := 1, 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			cnt++
		} else {
			ret := getCnt(cnt)
			copy(chars[ans:ans+len(ret)], ret)
			ans += len(ret) + 1
			cnt = 1
			chars[ans-1] = chars[i]
		}
	}
	ret := getCnt(cnt)
	copy(chars[ans:ans+len(ret)], ret)
	return ans + len(ret)
}

//leetcode submit region end(Prohibit modification and deletion)
