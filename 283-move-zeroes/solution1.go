func moveZeroes(nums []int) {
	n := len(nums)
	zeroCount := 0
	for i := 0; i < n-zeroCount; i++ {
		if nums[i] == 0 {
			for j := i; j < n-zeroCount-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			zeroCount++
			if nums[i] == 0 {
				i--
			}
		}
	}
}
