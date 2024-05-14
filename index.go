package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	lS := 0
	for range s {
		lS++
	}
	lSs := 0
	for range toFind {
		lSs++
	}

	if lS < lSs {
		return -1
	}

	indexOfSs := -1
	isMatch := false

	for i := 0; i < lS; i++ {
		if s[i] == toFind[0] {

			for j := 0; j < lSs; j++ {
				if i+j < lS && (s[i+j] == toFind[j]) {
					indexOfSs = i
					isMatch = true
				} else {
					indexOfSs = -1
					isMatch = false
				}
			}

			if isMatch {
				break
			}
		}
	}

	return indexOfSs
}
