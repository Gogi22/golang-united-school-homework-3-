package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	digits := make([]int, 0, len(input))

	for k := range input {
		digits = append(digits, k)
	}

	sort.Ints(digits)

	for _, k := range digits {
		result = append(result, input[k])
	}

	return
}
