package main

import "fmt"

func main() {
	words := []string{"a", "b", "c", "d"}
	groups := []int{1, 0, 1, 1}
	n := 4
	subsequence := getWordsInLongestSubsequence(n, words, groups)
	fmt.Println(subsequence)
}
func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	ret := []string{words[0]}
	for i := 1; i < n; i++ {
		if groups[i] == groups[i-1] {
			continue
		}
		ret = append(ret, words[i])
	}
	return ret
}
