package main

import (
	"fmt"
)

type Wallet struct{
	Cash int
}

func main(){
	myWallet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment : %#v\n", myWallet)
	fmt.Printf("Способ оплаты : %s\n", myWallet)
}
