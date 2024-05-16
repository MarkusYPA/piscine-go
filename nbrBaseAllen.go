package piscine

func NbrBaseAllen(nbr int, base string) string {
	baseLen := len(base)

	isNegative := false
	if nbr < 0 {
		isNegative = true
		if nbr == -9223372036854775808 {
			return "-9223372036854775808" // Avoids error but thats it?
		} else {
			nbr = -nbr
		}
	}

	newNbr := ""
	for nbr > 0 {
		newNbr = string(base[nbr%baseLen]) + newNbr // Start from the end and add the last digit to the start of the string
		nbr = nbr / baseLen                         // Remove the last number from the int
	}

	if isNegative {
		newNbr = "-" + newNbr
	}
	return (newNbr)
}
