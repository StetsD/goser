package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	ID int `json:"user_id,string"`
	Username string
	Address string `json:",omitempty"`
	Company string `json:"-"`
}

func main(){
	u := &User{
		ID: 42,
		Username: "boris",
		Address: "",
		Company: "Zazaza",
	}
	result, _ := json.Marshal(u)
	fmt.Printf("json string: %s\n", string(result))
}
