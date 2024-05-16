package piscine

func SplitWhiteSpaces(s string) []string {
	str := []string{}
	str = append(str, "") // String to index 0
	runes := []rune(s)

	index := 0

	for i := 0; i < len(runes); i++ {
		toAdd := string(runes[i])

		if isWhite(runes[i]) { // ascii 9 is tab

			if len(runes) >= i+1 && !isWhite(runes[i+1]) {
				index++
				str = append(str, "") // New string to add to unless two white spaces
			}

			continue
		}

		str[index] += toAdd
	}

	return str
}

func isWhite(r rune) bool {
	if r == 9 || r == ' ' || r == '\n' { // ascii 9 is tab
		return true
	}
	return false
}
