package piscine

func Map(f func(int) bool, a []int) []bool {
	bools := []bool{}
	for _, nbr := range a {
		bools = append(bools, f(nbr))
	}
	return bools
}
