package piscine

func IsUpper(s string) bool {
	allUppers := true
	for _, v := range s {
		if v > 'A' && v < 'Z' {
			allUppers = true
		} else {
			allUppers = false
		}
	}
	return allUppers
}
