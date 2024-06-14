package main

func isValid(s string) bool {
	st := []int32{0}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			st = append(st, c)
			continue
		case ')':
			if st[len(st)-1] != '(' {
				return false
			}
		case ']':
			if st[len(st)-1] != '[' {
				return false
			}
		case '}':
			if st[len(st)-1] != '{' {
				return false
			}
		}
		st = st[:len(st)-1]
	}
	return len(st) == 1
}
