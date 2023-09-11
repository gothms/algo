package main

import "fmt"

func main() {
	s1 := "abcdba"
	s2 := "cabdab"
	strings := checkStrings(s1, s2)
	fmt.Println(strings)
}
func checkStrings(s1 string, s2 string) bool {
	n := len(s1)
	cache := [26][2]byte{}
	for i := 1; i < n; i += 2 {
		cache[s1[i-1]-'a'][0]++
		cache[s2[i-1]-'a'][0]--
		cache[s1[i]-'a'][1]++
		cache[s2[i]-'a'][1]--
	}
	if n&1 == 1 {
		cache[s1[n-1]-'a'][0]++
		cache[s2[n-1]-'a'][0]--
	}
	for _, c := range cache {
		if c[0] != 0 || c[1] != 0 {
			return false
		}
	}
	return true
}
