import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res [][]int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			threeSum := v + nums[left] + nums[right]
			if threeSum > 0 {
				right--
			} else if threeSum < 0 {
				left++
			} else {
				res = append(res, []int{v, nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return res
}
