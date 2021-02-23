func isValid(s string) bool {
	const (
		OPEN_ROUND   = byte('(')
		OPEN_FIGURE  = byte('{')
		OPEN_SQUARE  = byte('[')
		CLOSE_ROUND  = byte(')')
		CLOSE_FIGURE = byte('}')
		CLOSE_SQUARE = byte(']')
	)
	brackets := map[byte]byte{
		CLOSE_ROUND:  OPEN_ROUND,
		CLOSE_FIGURE: OPEN_FIGURE,
		CLOSE_SQUARE: OPEN_SQUARE,
	}

	v := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == OPEN_ROUND || s[i] == OPEN_FIGURE || s[i] == OPEN_SQUARE {
			v = append(v, s[i])
		} else if len(v) > 0 && (s[i] == CLOSE_ROUND || s[i] == CLOSE_FIGURE || s[i] == CLOSE_SQUARE) {
			if v[len(v)-1] != brackets[s[i]] {
				return false
			}
			v = v[:len(v)-1]
		} else {
			return false
		}
	}

	return len(v) == 0
}
