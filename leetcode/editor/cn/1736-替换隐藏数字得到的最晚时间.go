package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumTime(time string) string {
	// lc
	var q byte = '?'
	t := []byte(time)
	if t[0] == q {
		if t[1] == q || t[1] <= '3' {
			t[0] = '2'
		} else {
			t[0] = '1'
		}
	}
	if t[1] == q {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == q {
		t[3] = '5'
	}
	if t[4] == q {
		t[4] = '9'
	}
	return string(t)

	// 个人
	//var sb strings.Builder
	//var q uint8 = '?'
	//if time[0] == q {
	//	if time[1] == q || time[1]-'0' <= 3 {
	//		sb.WriteRune('2')
	//	} else {
	//		sb.WriteRune('1')
	//	}
	//} else {
	//	sb.WriteByte(time[0])
	//}
	//if time[1] == q {
	//	if v := sb.String()[0] - '0'; v <= 1 {
	//		sb.WriteRune('9')
	//	} else {
	//		sb.WriteRune('3')
	//	}
	//} else {
	//	sb.WriteByte(time[1])
	//}
	//sb.WriteByte(time[2])
	//if time[3] == q {
	//	sb.WriteRune('5')
	//} else {
	//	sb.WriteByte(time[3])
	//}
	//if time[4] == q {
	//	sb.WriteRune('9')
	//} else {
	//	sb.WriteByte(time[4])
	//}
	//return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
