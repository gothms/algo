package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []Order{{3, 10}, {2, 6}, {5, 7}, {1, 9}, {6, 3}}
	sortOrders(arr, len(arr))
	fmt.Println(arr)
}

type Order struct {
	createTime int
	price      int
}

func sortOrders(orders []Order, n int) {
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].createTime < orders[j].createTime ||
			orders[i].createTime == orders[j].createTime && orders[i].price < orders[j].price
	})
}
