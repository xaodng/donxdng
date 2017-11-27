package main

import "fmt"

//返回函数
func create_func() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

//参数为函数
func hand(myfunc func(x,y int) int,a,b int) int {
	return myfunc(a,b)
}

func add(x,y int) int {
	return x + y
}

func main()  {
	func1 := create_func() //返回值为函数
	fmt.Println(func1(1,10))

	fmt.Println(hand(add, 100, 1000)) //传入参数为函数
}
