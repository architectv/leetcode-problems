import "strings"

func isPalindrome(s string) bool {
	var b strings.Builder
	for _, v := range strings.ToLower(s) {
		if v >= 'a' && v <= 'z' || v >= '0' && v <= '9' {
			b.WriteRune(v)
		}
	}

	ss := b.String()
	n := len(ss)
	for i := 0; i < n/2; i++ {
		if ss[i] != ss[n-1-i] {
			return false
		}
	}

	return true
}
