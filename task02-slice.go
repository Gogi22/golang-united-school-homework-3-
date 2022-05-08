package homework

func reverse(input []int64) (result []int64) {
	//Place your code here

	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-i-1] = input[len(input)-i-1], input[i]
	}

	result = input
	return
}
