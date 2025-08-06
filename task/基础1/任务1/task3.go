package main

/*
题目：查找字符串数组中的最长公共前缀
*/

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func findMinLengthStr(vs []string) string {
	var minIdx int
	minLength := utf8.RuneCountInString(vs[0])
	for idx, v := range vs {
		if idx == 0 {
			continue
		}
		length := utf8.RuneCountInString(v)
		if length < minLength {
			minLength = length
			minIdx = idx
		}
	}

	return vs[minIdx]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	dest := findMinLengthStr(strs)

	var prefix bytes.Buffer
	for idx, val := range dest {
		for _, item := range strs {
			temp := item[idx]
			if int32(temp) != val {
				return prefix.String()
			}
		}

		prefix.WriteRune(val)

	}
	return prefix.String()
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
