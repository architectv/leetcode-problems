func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	nums := make([]bool, n)
	count := 0

	for i := 2; i < n; i++ {
		if !nums[i] {
			count++
			for j := 2 * i; j < n; j += i {
				nums[j] = true
			}
		}
	}

	return count
}