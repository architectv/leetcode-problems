func merge(nums1 []int, m int, nums2 []int, n int) {
	k := n + m - 1

	for n > 0 && m > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[k] = nums1[m-1]
			m--
		} else {
			nums1[k] = nums2[n-1]
			n--
		}
		k--
	}

	for n > 0 {
		nums1[k] = nums2[n-1]
		n--
		k--
	}
}
