package main

import (
	"fmt"
	"encoding/json"
)

var jsonStr = `[
	{"id": 17, "username": "boris", "phone": 0},
	{"id": "17", "address": "none", "company": "Alalala"}
]`

func main(){
	data := []byte(jsonStr)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("unpacked in empty interface:\n%#v\n\n", user1)

	user2 := map[string]interface{}{
		"id": 42,
		"username": "Petr",
	}

	var user2i interface{} = user2
	result, _ := json.Marshal(user2i)
	fmt.Printf("json string from map:\n %s\n", string(result))
}
