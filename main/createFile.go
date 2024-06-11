package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const pkg, notes = "package ", "\n/*\n\n\n\n\n*/"
	const prePath = "math/"
	root, _ := os.Getwd()
	files := []string{"01.basic", "02.probability_theory", "03.linear_algebra", "04.practical", "05.additional"}
	ids := [][2]int{{1, 18}, {19, 32}, {33, 45}, {46, 51}, {1, 3}}
	const tag = false // Tag
	if !tag {
		return
	}
	for i, file := range files {
		os.Mkdir(prePath+file, os.ModePerm)
		str := strings.Split(file, ".")
		if len(str) < 2 { // 没数据
			break
		}
		id, _ := strconv.Atoi(str[0])
		pkgName := fmt.Sprintf("_%d_%s\n", id, str[1])
		for j := ids[i][0]; j <= ids[i][1]; j++ {
			fileName := fmt.Sprintf("%.2d", j)
			f, _ := os.Create(fmt.Sprintf("%s/%s/%s..go", root, prePath+file, fileName))
			os.WriteFile(f.Name(), []byte(pkg+pkgName+notes), 0)
		}
	}
}
