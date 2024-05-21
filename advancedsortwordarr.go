package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	// Standard bubble sort
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			// If f(a,b) returns a positive number (like Compare when a > b), switch places
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
