package main

import "fmt"

func main(){
	defer fmt.Println("After work")
	fmt.Println("Some userful work")
}
