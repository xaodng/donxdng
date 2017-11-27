package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var (
		name string
		age float64
		f interface{}
	)
	b := `{"Name":"Wednesday",
	"Age":6,
	"Parents":["Gomez","Morticia"]}`

	err := json.Unmarshal([]byte(b),&f)
	if err != nil{
		fmt.Println(err)
	}

	m := f.(map[string]interface{})

	for k, v := range m{
		switch vv := v.(type) {
		case string:
			fmt.Println(k ," is string ",vv)
			name = vv
		case float64:
			fmt.Println(k ," is int ",vv)
			age = vv
		case []interface{}:
			fmt.Println(k ," is an array: ")
			for i,u := range vv{
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	fmt.Println("[",name,"][",age,"]")
}
