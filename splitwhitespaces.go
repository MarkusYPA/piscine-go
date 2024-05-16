package piscine

func SplitWhiteSpaces(s string) []string {
	str := []string{}
	str = append(str, "")

	index := 0

	for i, v := range s {
		if v == 9 || v == ' ' || (len(s) > i+1 && s[i:i+2] == "\n") { // ascii 9 is tab
			index++
			str = append(str, "")
			if s[i:i+2] == "\n" {
				i++
			}
			continue
		}
		str[index] += string(v)
	}

	return str
}
