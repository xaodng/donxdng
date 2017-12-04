package main

import (
	"math"
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Sharper2 interface {
	Area2() float32
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func (ci Circle) Area2() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main()  {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	//注意是*Square，而不能是Square// 检查areaIntf是不是指向*Square类型
	if t,ok := areaIntf.(*Square); ok {
		fmt.Printf("areaIntf is; %T | %#v\n", t, t)
	}

	if u,ok := areaIntf.(*Circle);ok {
		fmt.Println("areaIntf is %T | %v\n", u, u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	//-----------------------------------------------

	sq2 := Circle{radius:1.3}
	var areaIntf2 Sharper2 = sq2
	if t, ok := areaIntf2.(Circle); ok {
		fmt.Printf("areaIntf2 is : %t=T | %#v\n", t, t)
	}

	 areaIntf2  = &sq2
	if t, ok := areaIntf2.(*Circle); ok {
		fmt.Printf("areaIntf2 is : %t=T | %#v\n", t, t)
	}

	fmt.Println("END")
}
