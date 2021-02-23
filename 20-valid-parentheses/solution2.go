func isValid(s string) bool {
	const (
		OPEN_ROUND   = byte('(')
		OPEN_FIGURE  = byte('{')
		OPEN_SQUARE  = byte('[')
		CLOSE_ROUND  = byte(')')
		CLOSE_FIGURE = byte('}')
		CLOSE_SQUARE = byte(']')
	)
	mapping := map[byte]byte{
		CLOSE_ROUND:  OPEN_ROUND,
		CLOSE_FIGURE: OPEN_FIGURE,
		CLOSE_SQUARE: OPEN_SQUARE,
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if v, ok := mapping[s[i]]; ok {
			topElement := byte('#')
			if len(stack) > 0 {
				n := len(stack) - 1
				topElement = stack[n]
				stack = stack[:n]
			}

			if v != topElement {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
