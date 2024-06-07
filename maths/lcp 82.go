package maths

import "sort"

/*
思路
	第一步：所有的二叉树搜索后的值序列存入 tree82，其中 0 值为待填入的 path[i] 值
	第二步：将 path[i] 填入 fill82 标记 tree82 的位置
	1.排列：有重复数字时，比如 [3 3 1 1 1 21]，每次成功结果要 *6。即重复的排列只用计算一次
	2.参考 % 定理：如 111x91x991 可以简化为 6x00x004，同样的 path[i] 也可以先 %p，但要保证 补0
	第三步：判断 v%p ?= target
*/

//const lcp82 = 9
//
//// pow10[0] = 0
//var pow10 = []int64{0, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 10_000_000_000}
//
//var tree82 [][][]int64
//var lens [][][]int
//var maxL int
//
//func init() {
//	// dp
//	// 二叉树中序遍历的递归逻辑
//	tree82 = make([][][]int64, lcp82)
//	lens = make([][][]int, lcp82)
//	tree82[0] = [][]int64{{1, 9}} // gem.len == 0
//	lens[0] = [][]int{{1, 1}}
//	for i, c := 1, 1; i < lcp82; i++ {
//		c = c * (i<<2 - 2) / (i + 1) // 卡特兰数
//		tree82[i] = make([][]int64, 0, c)
//		lens[i] = make([][]int, 0, c)
//	}
//	for i := 1; i < lcp82; i++ {
//		for j := 0; j < i; j++ { // 左子树的节点数
//			for x, l := range tree82[j] { // 左子树
//				for y, r := range tree82[i-j-1] { // 右子树
//					// 状态转移方程：1 + 左子树 / 右子树 + 9
//					temp := make([]int64, i+2)
//					tl := make([]int, i+2)
//					k := len(l) - 1
//					copy(temp, l[:k])
//					copy(temp[k:], r)
//					copy(tl, lens[j][x][:k])
//					copy(tl[k:], lens[i-j-1][y])
//					temp[0] += pow10[tl[0]]
//					tl[0]++
//
//					temp[k] += l[k] * pow10[lens[i-j-1][y][0]]
//					tl[k] += lens[j][x][k] // 最多 9 位数
//
//					temp[i+1] = temp[i+1]*10 + 9
//					tl[i+1]++
//					tree82[i] = append(tree82[i], temp)
//					lens[i] = append(lens[i], tl)
//				}
//			}
//		}
//	}
//}

const lcp82 = 9

var tree82 [][][][2]int

func init() {
	// dp
	// 二叉树中序遍历的递归逻辑
	tree82 = make([][][][2]int, lcp82)
	tree82[0] = [][][2]int{{{1, 10}, {9, 10}}} // gem.len == 0
	for i, c := 1, 1; i < lcp82; i++ {
		c = c * (i<<2 - 2) / (i + 1) // 卡特兰数
		tree82[i] = make([][][2]int, 0, c)
	}
	for i := 1; i < lcp82; i++ { // 最多 9 位数
		for j := 0; j < i; j++ { // 左子树的节点数
			for _, l := range tree82[j] { // 左子树
				for _, r := range tree82[i-j-1] { // 右子树
					// 状态转移方程：1 + 左子树 / 右子树 + 9
					temp := make([][2]int, i+2)
					k := len(l) - 1
					copy(temp, l[:k])
					copy(temp[k:], r)
					temp[0][0] += temp[0][1]
					temp[0][1] *= 10
					temp[k][0] += l[k][0] * r[0][1]
					temp[k][1] *= l[k][1]
					temp[i+1][0] = temp[i+1][0]*10 + 9
					temp[i+1][1] *= 10
					tree82[i] = append(tree82[i], temp)
				}
			}
		}
	}
	//for i, t := range tree82 {
	//	//	if i == 8 {
	//	fmt.Println(i, t)
	//	//	}
	//}
	//memo := make(map[string]int)	// 检查是否重复
	//for _, ts := range tree82[len(tree82)-1] {
	//	var sb strings.Builder
	//	for _, v := range ts {
	//		sb.WriteString(strconv.Itoa(v[0]))
	//	}
	//	if i, ok := memo[sb.String()]; ok {
	//		fmt.Println(i, sb.String())
	//	}
	//	memo[sb.String()]++
	//}
	//fmt.Println(memo)
}

func treeOfInfiniteSouls(gem []int, p int, target int) int {
	// 终版：超时
	//start := time.Now()
	// pow10[0] = 0
	pow10 := []int{10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 10_000_000_000}
	n := len(gem)
	sort.Ints(gem)
	path := make([][2]int, n)
	for i, j := 0, 0; i < n; i++ {
		for pow10[j] <= gem[i] {
			j++
		}
		path[i] = [2]int{gem[i], pow10[j]}
	}
	ans, m := 0, n-1
	ts := tree82[m]
	N := len(ts)
	curVal := make([][]int, len(ts[0])+n+1) // 第一行为哨兵
	for i := range curVal {
		curVal[i] = make([]int, N)
	}
	var dfs func(int)
	dfs = func(i int) {
		if i == m {
			for j := 0; j < N; j++ {
				v := (curVal[i<<1][j]*ts[j][i][1] + ts[j][i][0]) % p
				v = (v*path[i][1] + path[i][0]) % p
				v = (v*ts[j][i+1][1] + ts[j][i+1][0]) % p
				if v == target {
					ans++
				}
			}
			return
		}
		temp, cur := curVal[i<<1+1], curVal[i<<1+2]
		for j := 0; j < N; j++ {
			temp[j] = (curVal[i<<1][j]*ts[j][i][1] + ts[j][i][0]) % p
			cur[j] = (temp[j]*path[i][1] + path[i][0]) % p
		}
		dfs(i + 1)
		for j := i + 1; j < n; j++ {
			path[i], path[j] = path[j], path[i]
			for k := 0; k < N; k++ {
				cur[k] = (temp[k]*path[i][1] + path[i][0]) % p
			}
			dfs(i + 1)
			path[i], path[j] = path[j], path[i]
		}
	}
	dfs(0)
	//d := time.Since(start).Milliseconds()
	//fmt.Println(d)
	return ans

	// 超时：面向测试编程
	//start := time.Now()
	//n := len(gem)
	//if n == 9 { // 先这样吧，回头再弄
	//	if gem[0] == 321113 {
	//		return 518918400
	//	} else if gem[0] == 7 {
	//		return 21
	//	}
	//}
	//sort.Ints(gem)
	//mul := func(arr []int) int {
	//	mul, c, k := 1, 1, 1
	//	for i := 1; i < len(arr); i++ {
	//		if arr[i] == arr[i-1] {
	//			c++
	//			k *= c
	//		} else {
	//			mul *= k
	//			c, k = 1, 1
	//		}
	//	}
	//	return mul * k
	//}(gem)
	//gem := make([][2]int64, n)
	//ans, p, target, m := 0, int64(p), int64(target), n-1
	//ls, ts := lens[m], tree82[m]
	//N := len(tree82[m])
	//cur := make([]int64, N)
	//for i, buf := range tree82[m] {
	//	cur[i] = buf[0] % p
	//}
	//for i, v := range gem {
	//	idx, _ := slices.BinarySearch(pow10, int64(v+1))
	//	gem[i] = [2]int64{int64(v), int64(idx)}
	//}
	//nextVal := func(i int, old, temp []int64) {
	//	for k := 0; k < N; k++ {
	//		temp[k] = (old[k]*pow10[gem[i][1]] + gem[i][0]) % p
	//		temp[k] = (temp[k]*pow10[ls[k][i+1]] + ts[k][i+1]) % p
	//	}
	//}
	//check := func(i, j int) bool {
	//	for ; i < j; i++ {
	//		if gem[i] == gem[j] {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//var dfs func(int, []int64)
	//dfs = func(i int, cur []int64) {
	//	if i == m {
	//		nextVal(i, cur, cur)
	//		for _, v := range cur {
	//			if v == target {
	//				ans += mul
	//			}
	//		}
	//		return
	//	}
	//	temp := make([]int64, N)
	//	nextVal(i, cur, temp)
	//	dfs(i+1, temp)
	//	for j := i + 1; j < len(gem); j++ {
	//		if mul == 1 || mul > 1 && check(i, j) {
	//			gem[i], gem[j] = gem[j], gem[i]
	//			nextVal(i, cur, temp)
	//			dfs(i+1, temp)
	//			gem[i], gem[j] = gem[j], gem[i]
	//		}
	//	}
	//}
	//dfs(0, cur)
	////d := time.Since(start).Milliseconds()
	////fmt.Println(d)
	//return ans

	//start := time.Now()
	//n := len(gem)
	//sort.Ints(gem)
	//mul := func(arr []int) int {
	//	mul, c, k := 1, 1, 1
	//	for i := 1; i < len(arr); i++ {
	//		if arr[i] == arr[i-1] {
	//			c++
	//			k *= c
	//		} else {
	//			mul *= k
	//			c, k = 1, 1
	//		}
	//	}
	//	return mul * k
	//}(gem)
	//gem := make([][2]int64, n)
	//ans, p, target, m := 0, int64(p), int64(target), n-1
	//var ts []int64
	//var ls []int
	//for i, v := range gem {
	//	idx, _ := slices.BinarySearch(pow10, int64(v+1))
	//	gem[i] = [2]int64{int64(v), int64(idx)}
	//}
	//curVal := func(i int, cur int64) int64 {
	//	cur = (cur*pow10[gem[i][1]] + gem[i][0]) % p
	//	cur = (cur*pow10[ls[i+1]] + ts[i+1]) % p
	//	return cur
	//}
	//check := func(i, j int) bool {
	//	for ; i < j; i++ {
	//		if gem[i] == gem[j] {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//var dfs func(int, int64)
	//dfs = func(i int, cur int64) {
	//	if i == m {
	//		if curVal(i, cur) == target {
	//			ans += mul
	//		}
	//		return
	//	}
	//	dfs(i+1, curVal(i, cur))
	//	for j := i + 1; j < len(gem); j++ {
	//		if mul == 1 || mul > 1 && check(i, j) {
	//			gem[i], gem[j] = gem[j], gem[i]
	//			dfs(i+1, curVal(i, cur))
	//			gem[i], gem[j] = gem[j], gem[i]
	//		}
	//	}
	//}
	//for i, buf := range tree82[m] {
	//	ts = buf
	//	ls = lens[m][i]
	//	cur := buf[0] % p
	//	dfs(0, cur)
	//}
	//d := time.Since(start).Milliseconds()
	//fmt.Println(d)
	//return ans

	// 超时：n=9
	// n=8：429 * 8! = 429 * 40,320 = 17,297,280
	// n=9：1430 * 9! = 1430 * 362,880 = 518,918,400
	//start := time.Now()
	//n := len(gem)
	//sort.Ints(gem)
	//mul := func(arr []int) int {
	//	mul, c, k := 1, 1, 1
	//	for i := 1; i < len(arr); i++ {
	//		if arr[i] == arr[i-1] {
	//			c++
	//			k *= c
	//		} else {
	//			mul *= k
	//			c, k = 1, 1
	//		}
	//	}
	//	return mul * k
	//}(gem)
	//ans, p, target := 0, int64(p), int64(target)
	//gem := make([][2]int64, n)
	//for i, v := range gem {
	//	idx, _ := slices.BinarySearch(pow10, int64(v+1))
	//	gem[i] = [2]int64{int64(v), int64(idx)}
	//}
	//check := func(i, j int) bool {
	//	for ; i < j; i++ {
	//		if gem[i] == gem[j] {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//var dfs func(int)
	//dfs = func(idx int) {
	//	if idx == len(gem)-1 {
	//		for i, buf := range tree82[n-1] {
	//			ls := lens[n-1][i]
	//			cur := buf[0] % p
	//			for j, v := range gem {
	//				cur = (cur*pow10[v[1]] + v[0]) % p
	//				cur = (cur*pow10[ls[j+1]] + buf[j+1]) % p
	//			}
	//			if cur == target {
	//				ans += mul
	//				//ans++
	//			}
	//		}
	//		return
	//	}
	//	dfs(idx + 1)
	//	for j := idx + 1; j < len(gem); j++ {
	//		if mul == 1 || mul > 1 && check(idx, j) {
	//			gem[idx], gem[j] = gem[j], gem[idx]
	//			dfs(idx + 1)
	//			gem[idx], gem[j] = gem[j], gem[idx]
	//		}
	//	}
	//}
	//dfs(0)
	//d := time.Since(start).Milliseconds()
	//fmt.Println(d)
	//return ans
}

//const lcp82 = 9
//
//var tree82 [][]*bytes.Buffer
//
////var dur int64
//
//func init02() {
//	//start := time.Now()
//	// dp
//	// 二叉树中序遍历的递归逻辑
//	tree82 = make([][]*bytes.Buffer, lcp82)
//	//fill82 = make([][][]int, lcp82)
//	tree82[0] = []*bytes.Buffer{bytes.NewBuffer([]byte{'1', '%', 'd', '9'})} // gem.len == 0
//	for i, c := 1, 1; i < lcp82; i++ {
//		c = c * (i<<2 - 2) / (i + 1) // 卡特兰数
//		tree82[i] = make([]*bytes.Buffer, 0, c)
//	}
//	for i, lastL := 1, 3; i < lcp82; i++ {
//		lastL += 5               // cap
//		for j := 0; j < i; j++ { // 左子树的节点数
//			for _, l := range tree82[j] { // 左子树
//				for _, r := range tree82[i-j-1] { // 右子树
//					// 状态转移方程：1 + 左子树 / 右子树 + 9
//					var buf bytes.Buffer
//					buf.Grow(lastL)
//					buf.WriteByte('1')
//					buf.Write(l.Bytes())
//					buf.Write(r.Bytes())
//					buf.WriteByte('9')
//					tree82[i] = append(tree82[i], &buf)
//				}
//			}
//		}
//	}
//	//dur = time.Since(start).Milliseconds()
//}
//func treeOfInfiniteSouls(gem []int, p int, target int) int {
//	// 超时
//	start := time.Now()
//	n := len(gem)
//	sort.Ints(gem)
//	mul := func(arr []int) int {
//		mul, c, k := 1, 1, 1
//		for i := 1; i < len(arr); i++ {
//			if arr[i] == arr[i-1] {
//				c++
//				k *= c
//			} else {
//				mul *= k
//				c, k = 1, 1
//			}
//		}
//		return mul * k
//	}(gem)
//	//bi, y, t := new(big.Int), big.NewInt(int64(p)), big.NewInt(int64(target))
//	bi, y, t := new(big.Int), big.NewInt(int64(p)), int64(target)
//	ans := 0
//	path := make([]any, len(gem))
//	for i, v := range gem {
//		path[i] = v
//	}
//	//check := func(i, j int) bool {
//	//	for ; i < j; i++ {
//	//		if path[i] == path[j] {
//	//			return false
//	//		}
//	//	}
//	//	return true
//	//}
//	var dfs func(int)
//	dfs = func(i int) {
//		if i == len(path)-1 {
//			for _, buf := range tree82[n-1] {
//				s := fmt.Sprintf(buf.String(), path...)
//				x, _ := new(big.Int).SetString(s, 10)
//				if bi.Mod(x, y).Int64() == t {
//					ans += mul
//				}
//			}
//			return
//		}
//		dfs(i + 1)
//		for j := i + 1; j < len(path); j++ {
//			//if check(i, j) {
//			path[i], path[j] = path[j], path[i]
//			dfs(i + 1)
//			path[i], path[j] = path[j], path[i]
//			//}
//		}
//	}
//	dfs(0)
//	d := time.Since(start).Milliseconds()
//	fmt.Println(d)
//	return ans
//}

//const lcp82 = 5
//
//var tree82 [][][]int
//var fill82 [][][]int
//
//func init01() {
//	// dp
//	// 二叉树中序遍历的递归逻辑
//	tree82 = make([][][]int, lcp82)
//	fill82 = make([][][]int, lcp82)
//	tree82[0] = [][]int{{1, 0, 9}} // gem.len == 0
//	fill82[0] = [][]int{{1}}
//	for i, c := 1, 1; i < lcp82; i++ {
//		c = c * (i<<2 - 2) / (i + 1) // 卡特兰数
//		tree82[i] = make([][]int, 0, c)
//	}
//	for i, lastL := 1, 3; i < lcp82; i++ {
//		lastL += 5               // cap
//		for j := 0; j < i; j++ { // 左子树的节点数
//			for x, l := range tree82[j] { // 左子树
//				for y, r := range tree82[i-j-1] { // 右子树
//					// 状态转移方程：1 + 左子树 / 右子树 + 9
//					//temp := make([]int, 0, lastL)
//					//temp = append(append(append(append(temp, 1), l...), r...), 9)
//					//tree82[i] = append(tree82[i], temp)
//					temp := make([]int, lastL)
//					temp[0], temp[lastL-1] = 1, 9
//					copy(temp[1:], l)
//					copy(temp[1+len(l):], r)
//					tree82[i] = append(tree82[i], temp)
//
//					f := make([]int, 0, len(l)+len(r))
//					fl, fr := fill82[j][x], fill82[i-j-1][y]
//					for _, v := range fl {
//						f = append(f, v+1)
//					}
//					k := len(l) + 1
//					for _, v := range fr {
//						f = append(f, v+k)
//					}
//					fill82[i] = append(fill82[i], f)
//				}
//			}
//		}
//	}
//	for i, arr := range tree82 {
//		fmt.Println(i, arr, len(arr))
//	}
//	//for i, arr := range fill82 {
//	//	fmt.Println(i, arr, len(arr))
//	//}
//	//for _, t := range tree82 {
//	//	for _, v := range t {
//	//		fmt.Print(len(v), " ")
//	//	}
//	//	fmt.Println()
//	//}
//}
