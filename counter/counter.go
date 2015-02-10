package counter

func Counter(size int) func() int {

	arr := make([]int, size)

	// populate the values of the array
	for i, _ := range arr {
		arr[i] = i + 1
	}

	return func() int {
		sum := 0
		for i := range arr {
			sum = arr[i]
		}
		return sum
	}

}
