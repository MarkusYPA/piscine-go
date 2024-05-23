package piscine

func CollatzCountdown(start int) int {
	steps := 0
	return collatzHelper(start, steps)
}

func collatzHelper(start, steps int) int {
	if start == 1 {
		return steps
	}

	steps++

	if start%2 == 0 {
		return collatzHelper(start/2, steps)
	} else {
		return collatzHelper(start*3+1, steps)
	}
}
