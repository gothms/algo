package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for {
		m := (l + r) >> 1
		switch guess(m) {
		case -1:
			r = m - 1
		case 1:
			l = m + 1
		case 0:
			return m
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func guess(num int) int { return 0 }
