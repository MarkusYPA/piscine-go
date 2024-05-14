package piscine

func AlphaCount(s string) int {
	alphas := 0
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			alphas++
		}
	}
	return alphas
}
