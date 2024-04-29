package main

import (
	str "algo/string"
	"fmt"
)

func main() {
	// KMP
	s := "packagABABCABAAe strpackage strpackage str"
	p := "ABABCABAA"
	p = "ABABCABAB"
	p = "abababzabababa"
	//s = "aaaaaaaa"
	//p = "aaa"
	kmp := str.KmpInternet(s, p)
	fmt.Println("kmp:", kmp)

	beauty := str.KmpBeauty(s, p)
	fmt.Println("beauty:", beauty)

	lc := str.KmpLc(s, p)
	fmt.Println("lc:", lc)
}
