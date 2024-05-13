package piscine

// Store values on this map
var fibMemo = make(map[int]int)

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	// Return from list if already calculated. map[i] returns value AND bool
	if val, ok := fibMemo[index]; ok {
		return val
	}

	if index < 2 {
		fibMemo[index] = index // First store, then return
		return index
	}

	fibMemo[index] = Fibonacci(index-1) + Fibonacci(index-2) // First store, then return
	return fibMemo[index]
}
