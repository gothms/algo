package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func initIO() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	var result int
	var sign int = 1
	b, _ := reader.ReadByte()
	for b < '0' || b > '9' {
		if b == '-' {
			sign = -1
		}
		b, _ = reader.ReadByte()
	}
	for b >= '0' && b <= '9' {
		result = result*10 + int(b-'0')
		b, _ = reader.ReadByte()
	}
	return result * sign
}

func writeInt(x int) {
	writer.WriteString(strconv.Itoa(x))
	writer.WriteByte('\n')
}

func main() {
	initIO()
	defer writer.Flush()

	n := readInt()
	for i := 0; i < n; i++ {
		a := readInt()
		b := readInt()
		result := a + b
		writeInt(result)
	}
}
