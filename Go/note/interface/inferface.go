package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}
//Area method has pointer receiver,所以Shaper只能被赋值为*Square,而不能被赋值为Square
func (sq *Square) Area() float32  {
	return sq.side * sq.side
}

func main() {
	//内建函数 new 用来分配内存，它的第一个参数是一个类型，不是一个值，
	// 它的返回值是一个指向新分配类型零值的指针.
	sq1 := new(Square)
	sq1.side = 5

	//areaIntf := sq1
	//fmt.Printf("The square has area: %f\nareaIntf type:%v\n",areaIntf.Area(),areaIntf)

	//var areaIntf Shaper
	//areaIntf = sq1
	//fmt.Printf("The square has area: %f\nareaIntf type:%#v\n\n",areaIntf.Area(),areaIntf)

	areaIntf := Shaper(sq1)
	fmt.Printf("The square has area: %f\nareaIntf type:%#v\n\n",areaIntf.Area(),areaIntf)

	sq2 := &Square{2.0}
	areaIntf = sq2

	fmt.Println("END")
}