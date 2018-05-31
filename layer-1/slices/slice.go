package main

import "fmt"

func main(){
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)

	fmt.Println(
		buf0,
		buf1,
		buf2,
		buf3,
		buf4,
		buf5,
	)
}
