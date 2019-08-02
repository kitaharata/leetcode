func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	m := make([]int, n+1)
	m[0], m[1] = 1, 1
	for i := 2; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}
