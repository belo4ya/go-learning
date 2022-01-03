package main

import (
	"fmt"
)

type Person struct {
	Account
	Id   int
	Name string
}

type Account struct {
	Id int
}

type Interfacer interface {
}

func main() {
	person := Person{
		Id:      10,
		Name:    "Алексей",
		Account: Account{15},
	}

	if res, _ := func(a bool) (bool, error) { return a, nil }(true); res {
		fmt.Println(res)
	}

	if res, _ := func(a bool) (bool, error) { return a, nil }(false); res {
		fmt.Println(res)
	}

	fmt.Println(person.Id)
	fmt.Printf("%v\n", person)
	fmt.Printf("%#v\n", person)
}
