package main

import "fmt"

func main(){
	var a1 [3]int
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	a3 := [...]int{1,2,3}
	fmt.Println("a3", a3)
}
