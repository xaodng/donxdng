package main

import "fmt"

type cal struct {
	mul func(x,y int) int
	a,b int
}

func main()  {
	x := cal{
			mul: func(x, y int) int {
				return x * y
			},
			a:10,
			b:20,
	}

	fmt.Println(x.mul(x.a, x.b))
}
