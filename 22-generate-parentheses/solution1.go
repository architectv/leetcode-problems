func generateParenthesis(n int) []string {
	var ans []string
	backtrack(&ans, "", 0, 0, n)
	return ans
}

func backtrack(ans *[]string, s string, left, right, n int) {
	if len(s) == 2*n {
		*ans = append(*ans, s)
		return
	}

	if left < n {
		backtrack(ans, s+"(", left+1, right, n)
	}

	if right < left {
		backtrack(ans, s+")", left, right+1, n)
	}
}
