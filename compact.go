package piscine

func Compact(ptr *[]string) int {
	nonZeros := 0
	newSlice := []string{}
	for _, str := range *ptr {
		if str != "" {
			nonZeros++
			newSlice = append(newSlice, str)
		}
	}
	*ptr = newSlice
	return nonZeros
}
