package main

import "fmt"

type Person struct {
	Id int
	Name string
}

func (p Person) UpdateName(name string){
	p.Name = name
}

func (p *Person) SetName(name string){
	p.Name = name
}

type Account struct {
	Id int
	Name string
	Person
}

func main(){
	pers := Person{1, "Vasily"}
	pers.SetName("Vasiliy Romanov")
	// (&pers).SetName("Vasiliy Romanov")
	// fmt.Printf("update person: %#v\n", pers)

	var acc Account = Account{
		Id: 1,
		Name: "Boris",
		Person: Person{
			Id: 2,
			Name: "Vasily Ivanov",
		},
	}

	acc.SetName("romanov.vasiliy")

	fmt.Printf("%#v\n", acc)
}
