package main

import "fmt"

func Bar(v *int) {
	*v += 10
}

func Foo(s *[]int) {
	length := len(*s)

	for i := 0; i < length; i++ {
		(*s)[i] *= (*s)[i]
	}
}

func main() {
	var num = 10
	Bar(&num)
	fmt.Println(num)

	s := []int{1, 2, 3, 4, 5}
	Foo(&s)
	fmt.Println(s)
}
