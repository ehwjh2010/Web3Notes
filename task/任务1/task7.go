package main

import "fmt"

func twoSum(nums []int, target int) []int {
	valMap := make(map[int]int)

	for idx, num := range nums {
		other := target - num
		oIdx, exists := valMap[other]
		if !exists {
			valMap[num] = idx
			continue
		}

		return []int{idx, oIdx}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 3, 4}, 7))
}
