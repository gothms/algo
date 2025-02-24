package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

//type Employee struct {
//	Id           int
//	Importance   int
//	Subordinates []int
//}

func getImportance(employees []*Employee, id int) int {
	memo := make(map[int]*Employee)
	for _, e := range employees {
		memo[e.Id] = e
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		e := memo[i]
		ans := e.Importance
		for _, c := range e.Subordinates {
			ans += dfs(c)
		}
		return ans
	}
	return dfs(id)
}

//leetcode submit region end(Prohibit modification and deletion)
