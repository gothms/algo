package divide

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type point struct {
	x, y float64
}

var arrNP []*point
var idxNP int
var cntNP int

// var once sync.Once

func OverNP() {
	fmt.Println(cntNP)
}

func points() []*point { // 模拟点集
	idxNP ^= 1 // arr 用两次
	if idxNP == 0 {
		return arrNP
	}
	//if idxNP != 0 { // arr 用三次
	//	idxNP = (idxNP + 1) % 3
	//	return arr
	//}
	//idxNP++
	cntNP++

	//once.Do(func() {
	arrNP = arrNP[:0]
	n := 100
	randN := n * n
	for i := 0; i < n; i++ {
		x := float64(rand.Intn(randN))
		y := float64(rand.Intn(randN))
		arrNP = append(arrNP, &point{x, y})
	}
	//})
	return arrNP
}

/*
平面最近点对
	二维平面上有 n 个点，如何快速计算出两个距离最近的点对？
	https://oi-wiki.org/geometry/nearest-points/
*/

// ====================分治算法====================

// NearestPointsDivide 分治算法
func NearestPointsDivide(ps []*point) float64 {
	ps = points()

	ans, n := math.MaxFloat64, len(ps)
	sort.Slice(ps, func(i, j int) bool { // 按 x 排序
		return ps[i].x < ps[j].x || ps[i].x == ps[j].x && ps[i].y < ps[j].y
	})
	mergeSort := func(l, m, r int) { // 合并两个有序切片
		temp := make([]*point, r-l+1)
		j, k := m+1, 0
		for i := l; i <= m; i++ {
			for ; j <= r && ps[i].y > ps[j].y; j++ {
				temp[k] = ps[j]
				k++
			}
			temp[k] = ps[i]
			k++
		}
		copy(ps[l:], temp[:k])
	}
	var divide func(int, int)
	divide = func(l, r int) {
		if l+3 >= r { // 4 个点，暴力计算 distance，终止递归
			for i := l; i < r; i++ {
				for j := i + 1; j <= r; j++ {
					ans = min(ans, distance(ps[i], ps[j]))
				}
			}
			//sort.SliceStable(ps[l:r+1], func(i, j int) bool {
			sort.Slice(ps[l:r+1], func(i, j int) bool { // 按 y 排序
				return ps[l+i].y < ps[l+j].y
			})
			return
		}
		m := (l + r) >> 1
		mx := ps[m].x
		divide(l, m)
		divide(m+1, r)
		mergeSort(l, m, r)
		// x 排序：用于分治
		// y 排序：用于归并后 distance 的计算

		temp := make([]*point, 0)       // 缓冲区
		for i, k := l, 0; i <= r; i++ { // 以 ps[i] 为关注点
			if math.Abs(ps[i].x-mx) < ans {
				// x 坐标限制 (mx-ans, mx+ans)
				// y 坐标限制 (ps[i].y-ans, ps[i].y]，因为 [l,r] 区间已按 y 排序
				for j := k - 1; j >= 0 && ps[i].y-temp[j].y < ans; j-- {
					ans = min(ans, distance(ps[i], temp[j]))
				}
				temp = append(temp, ps[i]) // 分界点“周围” (mx-ans, mx+ans) 的点
				k++
			}
		}
	}

	divide(0, n-1)
	if ans == math.MaxFloat64 {
		return -1
	}
	return ans
}

// ====================有序数组====================

// NearestPointsNotDivide 非分治算法
// Red-Black Tree 实现数据有序
func NearestPointsNotDivide(ps []*point) float64 {
	ps = points()

	ans, n := math.MaxFloat64, len(ps)
	s := make([]*point, 0, n)
	//	fmt.Scan(&n)	// TODO
	//	for i := 0; i < n; i++ {
	//		fmt.Scan(&a[i].x, &a[i].y)
	//	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].x < ps[j].x || ps[i].x == ps[j].x && ps[i].y < ps[j].y
	})
	// 第 1、3 步都是 O(m) 的时间复杂度，最好是用 Red-Black Tree 实现 s 集合，则是 O(log m)
	for i, l := 0, 0; i < n; i++ {
		for l < i && ps[i].x-ps[l].x >= ans { // 1.将所有满足 xi-xl >= ans 的点从集合中删除，它们不会在对答案有贡献
			j := 0
			for j < len(s) && s[j] != ps[l] {
				j++
			}
			copy(s[j:], s[j+1:]) // 删除 ps[l]
			s = s[:len(s)-1]
			l++
		}

		j := sort.Search(len(s), func(idx int) bool { // 2.对于集合内满足 abs(yi-yj) < ans 的所有点，统计它们和 ps[i] 的距离
			return s[idx].y+ans > ps[i].y
		})
		for _, p := range s[j:] {
			ans = min(ans, distance(p, ps[i]))
		}

		s = append(s, ps[i]) // 3.将 ps[i] 插入到集合中，且保证按 y 有序
		last := len(s) - 2
		for last >= 0 && s[last].y > ps[i].y {
			s[last+1] = s[last]
			last--
		}
		s[last+1] = ps[i]
	}
	if ans == math.MaxFloat64 {
		return -1
	}
	return ans
}

// NearestPointsLine 期望线性做法
//func NearestPointsLine(ps []*point) float64 {
//	ps = points()
//
//	ans, n := math.MaxFloat64, len(ps)
//	sort.Slice(ps, func(i, j int) bool {
//		//return ps[i].x < ps[j].x || ps[i].x == ps[j].x && ps[i].y < ps[j].y
//		return ps[i].x < ps[j].x
//	})
//	if ans == math.MaxFloat64 {
//		return -1
//	}
//	return ans
//}

func distance(a, b *point) float64 {
	//return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
	return math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) // 优化：不求根
}

// ====================暴力====================

// NearestPointsForce 暴力法，用于验证其他方法
func NearestPointsForce(ps []*point) float64 {
	ps = points()

	ans, n := math.MaxFloat64, len(ps)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ans = min(ans, distance(ps[i], ps[j]))
		}
	}
	return ans
}
