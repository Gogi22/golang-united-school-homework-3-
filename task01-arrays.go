package homework

func average(input [15]float32) (result float32) {

	for _, v := range input {
		result += v
	}

	result /= float32(len(input))

	return
}
