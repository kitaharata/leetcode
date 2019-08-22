func reverse(x int) int {
	rev, pop, max, min := 0, 0, 1<<31-1, -1<<31
	for x != 0 {
		pop = x % 10
		x /= 10
		if rev > max/10 || rev == max/10 && pop > 7 {
			return 0
		}
		if rev < min/10 || rev == min/10 && pop < -8 {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
