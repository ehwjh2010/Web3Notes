package main

import "fmt"

/*
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度
*/
func removeDuplicates(nums []int) int {
	var uniqueIdx int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}

		if nums[i] != nums[uniqueIdx] {
			uniqueIdx++
			nums[uniqueIdx] = nums[i]
		}
	}

	return uniqueIdx + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
