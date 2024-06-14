package main

import (
	"fmt"
)

func main() {
	n, m := 10, 3 // [3,6,9,2,7,1,8,5,10,4]
	n, m = 10, 1  // [1,2,3,4,5,6,7,8,9,10]
	//n, m = 4, 5   // [1,3,4,2]
	i := joseph(n, m)
	fmt.Println(i)
}

//func joseph(n, m int) int {
//	if n == 1 {
//		return 1
//	}
//	return (joseph(n-1, m)+m-1)%n + 1
//}

// 数组
//func joseph(n int, m int) []int {
//	ans, arr := make([]int, n), make([]int, n)
//	for i := range arr {
//		arr[i] = i + 1
//	}
//	for i, j := 0, 0; i < n; i++ {
//		j = (j + m - 1) % len(arr)
//		ans[i] = arr[j]
//		arr = append(arr[:j], arr[j+1:]...)
//	}
//	return ans
//}

type ListNode struct {
	val  int
	next *ListNode
}

// 循环单链表
func joseph(n int, m int) []int {
	head := &ListNode{val: 1}
	cur := head
	for i := 2; i <= n; i++ {
		cur.next = &ListNode{val: i}
		cur = cur.next
	}
	pre := cur      // 前驱结点
	cur.next = head // 环
	m--
	ans := make([]int, 0, n)
	for i := n; i > 1; i-- { // TODO 优化：去掉 j 的 for 循环
		for j := 0; j < m%i; j++ { // 数 m-1 个结点
			pre = pre.next
		}
		ans = append(ans, pre.next.val)
		pre.next, pre.next.next = pre.next.next, nil // 删除结点
	}
	return append(ans, pre.val)
}
