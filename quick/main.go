package main

import "fmt"

func main(){
	nums := []int{5,9,7,4,1,2,3,11,13,6,2,5}

	//nums := []int{3,1,4,5,2,6}
	quickSort(nums)
	fmt.Println(nums)
}

func quickSort(input []int) []int{
	if len(input) < 2 {
		return input
	}

	st, ed := 0, len(input)-1
	input[st], input[ed] = input[ed], input[st] // pivot value 를 우측끝으로 보냄

	for idx:= range input{
		if input[idx] < input[ed] { // pivot 보다 작은값이면 맨앞부터 채워나감
			input[st], input[idx] = input[idx], input[st]
			st++
		}
	}

	input[st], input[ed] = input[ed], input[st]	// pivot 값을 가장 마지막 st 에다가 옮겨준다

	quickSort(input[:st])
	quickSort(input[st+1:])
	return input
}
