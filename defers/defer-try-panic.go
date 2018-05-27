package main

import (
	"fmt"
)

func deferTest(){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("try in panic")
			panic("uaaaa!")
		}
	}()
	fmt.Println("Some useful work")
	panic("Something bad happens")
	return
}

func main(){
	deferTest()
	return
}
