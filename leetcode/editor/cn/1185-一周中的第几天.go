//给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。
//
// 输入为三个整数：day、month 和 year，分别表示日、月、年。
//
// 您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday",
//"Friday", "Saturday"}。
//
//
//
// 示例 1：
//
// 输入：day = 31, month = 8, year = 2019
//输出："Saturday"
//
//
// 示例 2：
//
// 输入：day = 18, month = 7, year = 1999
//输出："Sunday"
//
//
// 示例 3：
//
// 输入：day = 15, month = 8, year = 1993
//输出："Sunday"
//
//
//
//
// 提示：
//
//
// 给出的日期一定是在 1971 到 2100 年之间的有效日期。
//
//
// Related Topics 数学 👍 130 👎 0

package main

import "fmt"

func main() {
	d, m, y := 1, 1, 1900
	d, m, y = 8, 1, 1900
	d, m, y = 1, 1, 1901
	d, m, y = 29, 2, 2016 // Monday
	d, m, y = 31, 8, 2019 // Saturday
	//d, m, y = 18, 7, 1999 // Sunday
	//d, m, y = 15, 8, 1993 // Sunday
	week := dayOfTheWeek(d, m, y)
	fmt.Println(week)
}

// leetcode submit region begin(Prohibit modification and deletion)
func dayOfTheWeek(day int, month int, year int) string {
	// 1900 年不是闰年，1904 年是闰年
	WEEK := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	md := [11]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	ds := (year-1900)*365 + (year-1901)/4 + day
	for i := 0; i < month-1; i++ {
		ds += md[i]
	}
	if month > 2 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		ds++
	}
	return WEEK[(ds+6)%7]
}

//leetcode submit region end(Prohibit modification and deletion)
