package oddsort

func oddSort(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)
	incr := 1

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			output[i+1] = input[i] + incr
			incr++
		} else {
			if output[i] >= input[i+1] {
				output[i+1] = input[i+1] + (incr - 1)
			} else {
				output[i+1] = input[i+1]
				incr = 1
			}
		}
	}
	return output
}

//input:  []int{1, 1, 1, 2, 7},
//wanted: []int{1, 2, 3, 4, 7},

func oddSortRef(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)
	incr := 1

	for i := 1; i < len(input); i++ {

		if input[i] == input[i-1] {
			output[i] = input[i-1] + incr
			incr++
		} else {
			if output[i-1] >= input[i] {
				output[i] = input[i] + (incr - 1)
			} else {
				output[i] = input[i]
				incr = 1
			}
		}
	}
	return output
}

func oddSort2(input []int) []int {
	output := make([]int, len(input))
	if len(input) == 0 {
		return output
	}
	output[0] = input[0]
	incr := 1

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			output[i+1] = input[i] + incr
			incr++
		} else {
			output[i+1] = input[i+1] + (incr - 1)
		}
	}
	return output
}
