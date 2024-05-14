package piscine

func Index(s string, toFind string) int {
	stR := []rune(s)
	sstR := []rune(toFind)

	lS := 0
	for range stR {
		lS++
	}
	lSs := 0
	for range sstR {
		lSs++
	}

	indexOfSs := -1
	isMatch := false

	for i := 0; i < lS; i++ {
		if stR[i] == sstR[0] {
			for j := 0; j < lSs; j++ {
				if stR[i+j] == sstR[j] {
					indexOfSs = i
					isMatch = true
				} else {
					indexOfSs = -1
				}
			}
			if isMatch {
				break
			}
		}
	}

	return indexOfSs
}
