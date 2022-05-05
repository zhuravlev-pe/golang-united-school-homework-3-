package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	for i, val := range input {
		result[len(input)-1-i] = val
	}
	return result
}
