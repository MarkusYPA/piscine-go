package piscine

func IsPrintable(s string) bool {
	notPrintable := false
	for _, v := range s {
		if v < ' ' {
			notPrintable = true
		}
	}
	return !notPrintable
}
