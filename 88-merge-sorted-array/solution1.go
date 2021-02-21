func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)
	i := 0
	j := 0
	k := 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res[k] = nums1[i]
			i++
		} else {
			res[k] = nums2[j]
			j++
		}
		k++
	}
	for i < m {
		res[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		res[k] = nums2[j]
		j++
		k++
	}
	for i := 0; i < m+n; i++ {
		nums1[i] = res[i]
	}
}
