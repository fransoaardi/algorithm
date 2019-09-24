package main

import "fmt"

func main(){
	//input := "한글(()깔깔(xx(xy(zxc)))"
	//input := ")한글 카카오 만(세)"
	//input := ")("
	input:= "{(x + 5) + 3} * y"
	//input := "({xx})"
	//input := "[{x()(x})]"

	//answer := intermediate(input)
	answer := hard(input)
	fmt.Println(answer)
}

func intermediate(input string) bool{
	stack := 0
	answer := false
	for _, char := range input{
		if string(char) == "(" {
			stack++
		}
		if string(char) == ")"{
			stack--
		}
		if stack < 0 {
			answer = false
			break
		}
	}
	if stack == 0 {
		answer = true
	}
	return answer
}

func hard(input string) bool{
	length := 1

	plusMap := map[string]int{
		"[": 1,		"{": 2,		"(": 3,
	}
	minusMap := map[string]int{
		"]": 1,		"}": 2,		")": 3,
	}

	stack := make([]int, len(input))
	answer := true

loop:
	for _, char := range input{
		if val, ok := plusMap[string(char)]; ok{
			if stack[length-1] < val {
				stack[length] = val
				length++
				//fmt.Println("a",stack, length)
			}else{
				answer = false
				break loop
			}
		}
		if val, ok := minusMap[string(char)]; ok{
			if stack[length-1] == val {
				//fmt.Println("b",stack, length)
				stack[length-1] = 0
				length--
			}else{
				answer = false
				break loop
			}
		}
	}
	return answer
}
