package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i, v := range strs {
		result += v
		if i < len(strs)-1 {
			result += sep
		}
	}
	return result
}
