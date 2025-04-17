package main

import "container/heap"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type TaskManager struct {
	m map[int][2]int // k:taskId,v:priority userId
	h *hp3408        // 大顶堆
}

func Constructor(tasks [][]int) TaskManager {
	m := make(map[int][2]int, len(tasks))
	h := &hp3408{}
	for _, t := range tasks {
		m[t[1]] = [2]int{t[2], t[0]}
		heap.Push(h, [2]int{t[2], t[1]})
	}
	return TaskManager{m, h}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.m[taskId] = [2]int{priority, userId}
	heap.Push(this.h, [2]int{priority, taskId})
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	p := this.m[taskId]
	p[0] = newPriority
	this.m[taskId] = p
	heap.Push(this.h, [2]int{newPriority, taskId})
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.m, taskId)
}

func (this *TaskManager) ExecTop() int {
	if len(this.m) == 0 {
		return -1
	}
	top := heap.Pop(this.h).([2]int)
	for this.m[top[1]][0] != top[0] {
		top = heap.Pop(this.h).([2]int)
	}
	user := this.m[top[1]][1]
	delete(this.m, top[1])
	return user
}

type hp3408 [][2]int // priority taskId

func (h hp3408) Len() int { return len(h) }
func (h hp3408) Less(i, j int) bool {
	return h[i][0] > h[j][0] || h[i][0] == h[j][0] && h[i][1] > h[j][1]
}
func (h hp3408) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp3408) Push(x any)   { *h = append(*h, x.([2]int)) }
func (h *hp3408) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
//leetcode submit region end(Prohibit modification and deletion)
