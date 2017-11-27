package main

import "fmt"

func main()  {
	p := Add()
	fmt.Printf("Call Add for 3 gives:%v\n",p(3)) //5

	p1 := Adder(2)
	fmt.Printf("The result is: %v\n",p1(3)) //5
	fmt.Printf("The result is: %v\n",p1(13)) //16
	fmt.Printf("The result is: %v\n",p1(13)) //17

	var f = Adder2()
	fmt.Print(f(1)," - ")
	fmt.Print(f(20)," - ")
	fmt.Print(f(300))
	fmt.Println()
	//1 - 21 - 321
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int  {
	return func(b int) int {
		defer func() {//证明不仅可以使用，还可以修改外部函数的变量
			a += 1
		}()
		return a + b
	}
}

func Adder2() func(int) int  {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}
