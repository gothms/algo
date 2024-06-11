package main

import (
	"fmt"
	"sort"
)

func main() {
	watchedVideos := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	friends := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}
	id, level := 0, 2
	byFriends := watchedVideosByFriends(watchedVideos, friends, id, level)
	fmt.Println(byFriends)
}

// leetcode submit region begin(Prohibit modification and deletion)
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {

}

//leetcode submit region end(Prohibit modification and deletion)

func watchedVideosByFriends_(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	// bfs
	n := len(friends)
	q, vis := []int{id}, make([]bool, n)
	vis[id] = true
	for i := 0; i < level && len(q) > 0; i++ {
		next := make([]int, 0)
		for _, j := range q {
			for _, k := range friends[j] {
				if vis[k] { // 重要：不能重复访问朋友
					continue
				}
				vis[k] = true
				next = append(next, k)
			}
		}
		q = next
	}
	memo := make(map[string]int) // 最终 level 的朋友
	for _, i := range q {
		for _, s := range watchedVideos[i] {
			memo[s]++ // 视频的频率
		}
	}
	ans := make([]string, 0, len(memo))
	for s := range memo {
		ans = append(ans, s)
	}
	sort.Slice(ans, func(i, j int) bool {
		return memo[ans[i]] < memo[ans[j]] || memo[ans[i]] == memo[ans[j]] && ans[i] < ans[j]
	})
	return ans
}
