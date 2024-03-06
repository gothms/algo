//给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。返回该日期是当年的第几天。
//
//
//
// 示例 1：
//
//
//输入：date = "2019-01-09"
//输出：9
//解释：给定日期是2019年的第九天。
//
// 示例 2：
//
//
//输入：date = "2019-02-10"
//输出：41
//
//
//
//
// 提示：
//
//
// date.length == 10
// date[4] == date[7] == '-'，其他的 date[i] 都是数字
// date 表示的范围从 1900 年 1 月 1 日至 2019 年 12 月 31 日
//
//
// Related Topics 数学 字符串 👍 123 👎 0

package main

import (
	"strconv"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func dayOfYear(date string) int {
	month := [12]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	y, _ := strconv.Atoi(date[:4])
	m, _ := strconv.Atoi(date[5:7])
	d, _ := strconv.Atoi(date[8:])
	for i := 0; i < m; i++ {
		d += month[i]
	}
	if m > 2 && (y%400 == 0 || y%4 == 0 && y%100 != 0) {
        d++
	}
	return d
}

//leetcode submit region end(Prohibit modification and deletion)
