package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	orderNow := 0
	firstOrder := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			orderNow = 1
		}
		if f(a[i], a[i+1]) < 0 {
			orderNow = -1
		}
		if orderNow == 1 || orderNow == -1 && firstOrder == 0 {
			firstOrder = orderNow
		}
		if orderNow == 1 || orderNow == -1 {
			if orderNow != firstOrder {
				return false
			}
		}
	}
	return true
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
