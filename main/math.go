package main

import (
	"algo/maths"
	"fmt"
	"math/rand"
)

func main() {
	a, mod := 15, 97
	a = rand.Intn(100) + 1
	//mod = rand.Intn(100) + 1
	i := maths.Inv(a, mod)
	fmt.Println(i)
	fmt.Println(maths.Gcd(a, mod) == 1, a*i%mod == 1)

	var x, y int = 99, 99
	maths.ExGcd(a, mod, &x, &y)
	fmt.Println(x, y)

	pow := maths.QPow(a, mod)
	//fmt.Println(pow)
	fmt.Println(a*((pow+mod)%mod)%mod == 1)

	mod = 16777619
	gcdN := maths.ExGcdN(mod-1, mod)
	for i, x := range gcdN {
		if (i+1)*x%mod != 1 {
			fmt.Println(i, x, mod, i*x%mod)
		}
	}

	a = 10
	p := 13
	n := maths.ExGcdN(a, p)
	fmt.Println(n)

	arr := []int{4, 7, 8, 12, 123456}
	p = 1_000_000_007
	s := maths.ExGcdS(arr, p)
	fmt.Println(s)
	//for _, v := range arr {
	//	maths.ExGcd(v, p, &x, &y)
	//	fmt.Print((x+p)%p, ",")
	//}
}
