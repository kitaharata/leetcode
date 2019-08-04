func lengthOfLastWord(s string) int {
	i := 0
	for j := range s {
		if s[j] != ' ' {
			if j > 0 && s[j-1] == ' ' {
				i = 0
			}
			i++
		}
	}
	return i
}
