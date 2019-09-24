package main

import "fmt"

func main(){
	var s, p string

	s = "BANANANANABA"
	p = "ANA"

	fail := failFunc(p)
	fmt.Println(fail)
	ans := kmp(s, p, fail)

	fmt.Println(ans)
}

func failFunc(p string)  []int {
	fail := make([]int, len(p))

	j := 0
	for i:=1; i<len(p); i++ {
		for j > 0 && p[i] != p[j] {
			//fmt.Println(i, j, string(p[i]), string(p[j]), fail[j-1])
			j = fail[j-1]
		}
		if p[i] == p[j]{
			j++
			fail[i] = j
		}
	}
	return fail
}

func kmp(s, p string, fail []int) (ans []int) {
	j := 0
	for i:=0; i<len(s); i++ {
		for j > 0 && s[i] != p[j] {
			j = fail[j-1]
		}
		if s[i] == p[j] {
			if j == len(p)-1{
				ans = append(ans, i-j)
				j = fail[j]
			}else{
				j++
			}
		}
	}
	return
}


