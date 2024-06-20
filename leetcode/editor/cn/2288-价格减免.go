package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	sentence := "1 2 $3 4 $5 $6 7 8$ $9 $10$"
	discount := 100 // 1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$
	sentence = "$2$3 $10 $100 $1 200 $33 33$ $$ $99 $99999 $9999999999"
	discount = 0 // $2$3 $10.00 $100.00 $1.00 200 $33.00 33$ $$ $99.00 $99999.00 $9999999999.00
	prices := discountPrices(sentence, discount)
	fmt.Println(prices)
}

// leetcode submit region begin(Prohibit modification and deletion)
func discountPrices(sentence string, discount int) string {
	n := len(sentence)
	var sb strings.Builder
	for i := 0; i < n; {
		sb.WriteByte(sentence[i])
		if sentence[i] == '$' && (i == 0 || sentence[i-1] == ' ') {
			i++
			j := i
			for i < n && unicode.IsDigit(rune(sentence[i])) {
				i++
			}
			if i > j && (i == n || sentence[i] == ' ') {
				v, _ := strconv.Atoi(sentence[j:i])
				sb.WriteString(fmt.Sprintf("%.2f", float64(v*(100-discount))/100))
			} else {
				sb.WriteString(sentence[j:i])
			}
		} else {
			i++
		}
	}
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
