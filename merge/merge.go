package main

import "fmt"

func main() {
	input := [][]int{{1, 3, 6, 10, 12}, {2, 8, 11}, {4, 7, 9}, {1, 2, 3, 4}, {3, 7, 9}}
	for len(input) > 1 {
		if len(input)%2 == 1 {
			input[len(input)-2] = merge(input[len(input)-2], input[len(input)-1])
			input = input[:len(input)-1]
		}
		for i := 0; i < len(input); i += 2 {
			input[i/2] = merge(input[i], input[i+1])
		}
		input = input[:len(input)/2]
	}
	fmt.Println(input[0])
}

func merge(input1, input2 []int) []int {
	rtn := make([]int, 0)
	for len(input1)*len(input2) != 0 {
		if input1[0] <= input2[0] {
			rtn = append(rtn, input1[0])
			input1 = input1[1:]
		} else {
			rtn = append(rtn, input2[0])
			input2 = input2[1:]
		}
	}
	if len(input1) != 0 {
		rtn = append(rtn, input1...)
	} else {
		rtn = append(rtn, input2...)
	}
	return rtn
}
