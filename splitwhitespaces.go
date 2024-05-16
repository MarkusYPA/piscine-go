package piscine

func SplitWhiteSpaces(s string) []string {
	str := []string{}
	str = append(str, "") // String to index 0

	index := 0

	for i := 0; i < len(s); i++ {
		toAdd := string(s[i])
		if s[i] == 9 || s[i] == ' ' || (len(s) > i+1 && s[i:i+2] == "\n") { // ascii 9 is tab
			index++
			str = append(str, "") // New string to add to
			if s[i:i+2] == "\n" {
				i++
			}
			toAdd = "" // Add empty when space, tab or newline

			if i == 0 && s[0] == ' ' {
				toAdd = " "
			}
		}

		str[index] += toAdd
	}

	return str
}
