package main

import (
	"fmt"
	"math"
)

func main() {
	n := 100
	sqrt := int(math.Ceil((-1 + math.Sqrt(float64(8*n+1))) / 2))
	//sqrt := math.Sqrt(float64(n << 3))
	v := sqrt * (1 + sqrt) >> 1
	fmt.Println(sqrt, v)

	sqrt = int(math.Sqrt(float64(2*n)+0.25) - 0.5)
	v = sqrt * (1 + sqrt) >> 1
	fmt.Println(sqrt, v)

	candies := distributeCandies(n, 4)
	fmt.Println(candies)
}

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(candies int, num_people int) []int {
	// lc
	n := num_people
	p := int(math.Sqrt(float64(2*candies)+0.25) - 0.5)
	rows, cols := p/n, p%n
	rowD := rows * (rows - 1) * n >> 1

	d := make([]int, n)
	for i := 0; i < n; i++ {
		// (i+1)*rows：i+1 为索引增量，每多一行则多一次增量
		// int(float64(rows*(rows-1)*n)*0.5)：每多一行则增量 n
		//d[i] = (i+1)*rows + int(float64(rows*(rows-1)*n)*0.5)
		d[i] = (i+1)*rows + rowD
		if i < cols {
			d[i] += i + 1 + rows*n
		}
	}
	d[cols] += candies - (p+1)*p>>1
	return d

	//ans := make([]int, num_people)
	//n := int(math.Ceil((-1 + math.Sqrt(float64(8*candies+1))) / 2))
	//for i := range ans {
	//	for j := i + 1; j <= n; j += num_people {
	//		ans[i] += j
	//	}
	//}
	//ans[(n-1)%num_people] -= n*(n+1)>>1 - candies
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
