package pkg

func Map(strs []string, mapFunc func(s string) string) []string {
	newStrs := make([]string, len(strs))
	for i, str := range strs {
		newStrs[i] = mapFunc(str)
	}
	return newStrs
}
