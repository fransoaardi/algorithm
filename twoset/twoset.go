package main

import (
	"fmt"
	"sort"
)

//예) [1, 3, 2, 4],  n = 5 -> 2 // (1,4), (2, 3)
func main(){
	n := 5
	ans := 0
	input := []int{1,3,2,4}

	sort.Ints(input)

	st, ed := 0, len(input)-1
	for st < ed {
		sum := input[st] + input[ed]
		switch{
		case sum<n:
			st++
			break
		case sum>n:
			ed--
			break
		case sum==n:
			fmt.Println(ans+1, ":", input[st], input[ed])
			st, ed = st+1, ed-1
			ans++
			break
		}
	}
	fmt.Println(ans, "sets")
}


