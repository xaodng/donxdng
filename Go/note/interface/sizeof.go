package main

import (
	"fmt"
	"unsafe"
)

type I interface {
}

func main()  {
	//str := "ssss"
	//var inter interface{} = str
	//fmt.Printf("%T	 | %#v\n", inter, inter)

	var i I
	fmt.Printf("%d, %v, %T\n",unsafe.Sizeof(i), i, i) //8, <nil>, <nil>

	s := "Hello, world!"
	i = s
	fmt.Printf("%d, %v, %T\n",unsafe.Sizeof(i), i, i)//8

	fmt.Println("END")
}
