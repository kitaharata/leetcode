func romanToInt(s string) int {
	res := 0
	value := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i, c := range s {
		if i != 0 && value[rune(s[i-1])] < value[c] {
			res -= 2 * value[rune(s[i-1])]
		}
		res += value[c]
	}
	return res
}
