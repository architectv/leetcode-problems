func groupAnagrams(strs []string) [][]string {
	const ALPHABET_LEN = 26
	dict := make(map[[ALPHABET_LEN]int][]string)

	for _, str := range strs {
		var count [ALPHABET_LEN]int
		for i := 0; i < len(str); i++ {
			idx := int(str[i]) - int(byte('a'))
			count[idx]++
		}
		dict[count] = append(dict[count], str)
	}

	groups := make([][]string, 0, len(dict))
	for alphabet := range dict {
		groups = append(groups, dict[alphabet])
	}

	return groups
}
