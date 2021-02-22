import "sort"

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	dict := make(map[string][]string, n)

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i int, j int) bool {
			return b[i] < b[j]
		})
		sortStr := string(b)
		dict[sortStr] = append(dict[sortStr], str)
	}

	groups := make([][]string, 0, len(dict))
	for sortStr := range dict {
		groups = append(groups, dict[sortStr])
	}

	return groups
}
