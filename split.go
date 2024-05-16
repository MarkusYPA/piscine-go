package piscine

func Split(s, sep string) []string {
	output := []string{}
	output = append(output, "") // String to index 0

	index := 0

	for i := 0; i < len(s); i++ {
		toAdd := s[i]

		// Check if at beginning of separator
		if len(s) >= i+len(sep) && s[i:i+len(sep)] == sep {

			// New string to add to unless next is separator too
			if len(s) >= i+len(sep)*2 && s[i+2:i+2+len(sep)] != sep {
				index++
				output = append(output, "")
			}

			i += len(sep) - 1
			continue
		}

		output[index] += string(toAdd)
	}

	return output
}
