package main

import (
	"fmt"
)

func main() {
	num := "123"
	target := 6
	num = "105"
	target = 5 // ["1*0+5","10-5"]
	num = "3456237490"
	target = 9191 // []
	operators := addOperators(num, target)
	fmt.Println(operators)
}

// leetcode submit region begin(Prohibit modification and deletion)
func addOperators(num string, target int) []string {
	// lc
	n := len(num)
	ans := make([]string, 0)
	path := make([]byte, 0, n<<1-1)
	var dfs func(int, int, int)
	dfs = func(i, ret, last int) {
		if i == n {
			if ret == target {
				ans = append(ans, string(path))
			}
			return
		}
		k := len(path)
		if i > 0 {
			path = append(path, 0) // 用于写入符号
		}
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ { // j == i || num[i] != '0'：不是前导 0
			path = append(path, num[j])
			val = val*10 + int(num[j]-'0')
			if i == 0 { // 填第一个数
				dfs(j+1, val, val)
			} else { // ret 当前结果，last 最后一个值
				path[k] = '+'
				dfs(j+1, ret+val, val)
				path[k] = '-'
				dfs(j+1, ret-val, -val)
				path[k] = '*'
				dfs(j+1, ret-last+last*val, last*val)
			}
		}
		path = path[:k] // 回溯
	}
	dfs(0, 0, 0)
	return ans

	// 个人
	//op := [3]byte{'+', '-', '*'} // 43 45 42 不包括除法
	//n := len(num)
	//path, option := make([]int, 0, n), make([]byte, 0, n-1)
	//getRet := func() int {
	//	vs, os := make([]int, 0, len(path)), make([]byte, 0, len(option))
	//	vs = append(vs, path[0])
	//	for i, o := range option { // 先算乘法
	//		v := path[i+1]
	//		if o == '*' {
	//			vs[len(vs)-1] *= v
	//		} else {
	//			vs, os = append(vs, v), append(os, o)
	//		}
	//	}
	//	ret := vs[0]
	//	for i, o := range os { // 易错四：减法运算时，符号在前，数字在后
	//		switch o {
	//		case '+':
	//			ret += vs[i+1]
	//		case '-':
	//			ret -= vs[i+1]
	//		}
	//	}
	//	return ret
	//}
	//ans := make([]string, 0)
	//var sb strings.Builder
	//var dfs func(int, int, int)
	//dfs = func(last, i, s int) {
	//	if i == n {
	//		if last == n-1 && getRet() == target { // 易错一：last == n-1，last 主要用于此处
	//			sb.Reset()
	//			for idx, o := range option {
	//				sb.WriteString(strconv.Itoa(path[idx]))
	//				sb.WriteByte(o)
	//			}
	//			sb.WriteString(strconv.Itoa(path[len(path)-1]))
	//			ans = append(ans, sb.String())
	//		}
	//		return
	//	}
	//	v := int(num[i] - '0')
	//	s = s*10 + v
	//	if v != 0 || last != i-1 { // 易错二：不能有前导 0
	//		dfs(last, i+1, s)
	//	}
	//	if len(path) == 0 { // 第一个数
	//		path = append(path, s)
	//		dfs(i, i+1, 0)
	//		path = path[:len(path)-1]
	//		return // 易错三：先有第一个数，才有后续运算
	//	}
	//	path = append(path, s)
	//	for _, o := range op {
	//		option = append(option, o)
	//		dfs(i, i+1, 0)
	//		option = option[:len(option)-1] // 回溯
	//	}
	//	path = path[:len(path)-1] // 回溯
	//}
	//dfs(-1, 0, 0)
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
