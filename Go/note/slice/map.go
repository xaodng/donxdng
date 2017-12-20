package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make (map[string]PersonInfo)

	personDB["111"] = PersonInfo{"12345","Tom", "Room203"}
	personDB["222"] = PersonInfo{"1", "Jack", "Room101"}

	person, ok := personDB["1111"]
	if ok {
		fmt.Println(person.Name, person.Address)
	}else {
		fmt.Println("Did not find person with ID 111")
	}
}
