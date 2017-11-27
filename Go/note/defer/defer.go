package main

import "fmt"

func f()(result int)  {
	defer func() {
		result++
	}()
	return 0
}

func f1()(r int)  {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func f2()(r int)  {
	defer func(i int) {
		i += 5
	}(r)
	return 2
}
//1. return xxx这条语句并不是一条原子指令，函数返回的过程是这样的：先给返回值赋值，然后调用 defer 表达式，最后才是返回到调用函数中。
func main()  {
	fmt.Println(f())

	fmt.Println(f1())

	fmt.Println(f2())
}