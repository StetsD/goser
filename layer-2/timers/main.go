package main

import (
	"time"
	"fmt"
)

func main(){
	timer := time.NewTimer(3 * time.Second)
	select {
	case <- timer.C:
		fmt.Println("timer.C timeout happend")
	case <- time.After(time.Minute):
		fmt.Println("time.After timeout happend")
	}
}
