package piscine

/* func ActiveBits(n int) int {
	binaryString := ToBase(n, "01")

	actives := 0
	for _, num := range binaryString {
		if num == '1' {
			actives++
		}
	}

	return actives
} */

func ActiveBits(n int) int {
	if n == -9223372036854775808 {
		return 1
	}

	if n < 0 {
		n *= -1
	}

	activeBits := 0
	for n > 0 {
		if n%2 == 1 {
			activeBits++
		}
		n /= 2
	}

	return activeBits
}
