import "strings"

func isPalindrome(s string) bool {
	ss := ""
	for _, v := range strings.ToLower(s) {
		if v >= 'a' && v <= 'z' || v >= '0' && v <= '9' {
			ss += string(v)
		}
	}

	n := len(ss)
	for i := 0; i < n/2; i++ {
		if ss[i] != ss[n-1-i] {
			return false
		}
	}

	return true
}
