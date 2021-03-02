func hammingWeight(num uint32) int {
	oneBits := 0
	for num > 0 {
		oneBits += int(num % 2)
		num /= 2
	}

	return oneBits
}
