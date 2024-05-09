package piscine

func StrRev(s string) string {
	sLen := stringLen(s)
	sRevRunes := make([]rune, sLen)

	for i, v := range s {
		sRevRunes[sLen-1-i] = rune(v)
	}

	sRevString := string(sRevRunes)

	return sRevString
}

func stringLen(s string) int {
	length := 0

	for range s {
		length++
	}

	return length
}
