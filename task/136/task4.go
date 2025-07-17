package main

import "fmt"

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(digits []int) []int {
	length := len(digits)

	if digits[length-1] != 9 {
		digits[length-1]++
		return digits
	}

	var which = length - 1
	for i := length - 2; i >= 0; i-- {
		if digits[i] == 9 {
			which = i
		} else {
			break
		}
	}

	var r []int
	if which == 0 {
		r = []int{1}
	} else {
		digits[which-1]++
		r = append(r, digits[:which]...)
	}

	right := make([]int, length-which)
	r = append(r, right...)
	return r
}

func main() {
	fmt.Println(plusOne([]int{9}))
}
