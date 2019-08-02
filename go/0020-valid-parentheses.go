func isValid(s string) bool {
	stack := make([]rune, len(s))
	i := 0
	for _, c := range s {
		switch c {
		case '(':
			stack[i] = ')'
			i++
		case '{':
			stack[i] = '}'
			i++
		case '[':
			stack[i] = ']'
			i++
		default:
			if i == 0 || stack[i-1] != c {
				return false
			}
			i--
		}
	}
	return i == 0
}
