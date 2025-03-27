package main

import (
	"fmt"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type timeVal struct {
	t   int
	val string
}
type TimeMap struct {
	db map[string][]timeVal
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]timeVal)}
}
func (this *TimeMap) search(key string, timestamp int) int {
	timestamp++
	if tv, ok := this.db[key]; ok {
		return sort.Search(len(tv), func(i int) bool {
			return tv[i].t >= timestamp
		}) - 1
	}
	return -1
}
func (this *TimeMap) Set(key string, value string, timestamp int) {
	//i := this.search(key, timestamp)
	//if i == -1 || this.db[key][i].t != timestamp {
	this.db[key] = append(this.db[key], timeVal{timestamp, value})
	//} else {
	//	this.db[key][i].val = value
	//}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := this.search(key, timestamp)
	if i == -1 {
		return ""
	} else {
		return this.db[key][i].val
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
//leetcode submit region end(Prohibit modification and deletion)
