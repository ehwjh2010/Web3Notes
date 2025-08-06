package main

import (
	"fmt"
	"sort"
)

/*
合并区间
*/

func mergeIntervals(vs [][]int) [][]int {
	if len(vs) <= 1 {
		return vs
	}

	sort.Slice(vs, func(i, j int) bool {
		return vs[i][0] < vs[j][0]
	})

	merged := [][]int{vs[0]}

	for _, v := range vs[1:] {
		length := len(merged)
		if v[0] > merged[length-1][1] {
			merged = append(merged, v)
		} else {
			if v[1] > merged[length-1][1] {
				merged[length-1][1] = v[1]
			}
		}
	}

	return merged
}

func main() {
	fmt.Println()
}
