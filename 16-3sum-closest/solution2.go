func threeSumClosest(nums []int, target int) int {
	diff := int(^uint(0) >> 1)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			complement := target - nums[i] - nums[j]
			hi := sort.Search(len(nums), func(k int) bool {
				if k <= j {
					return false
				}
				return nums[k] >= complement
			})
			lo := hi - 1
			if hi < len(nums) && abs(complement-nums[hi]) < abs(diff) {
				diff = complement - nums[hi]
			}
			if lo > j && abs(complement-nums[lo]) < abs(diff) {
				diff = complement - nums[lo]
			}
		}
		if diff == 0 {
			break
		}
	}

	return target - diff
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
