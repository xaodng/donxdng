package main

import "fmt"

func test1() []int {
	var prt []int
	arr := []int{1, 2, 3, 4, 5}

	for _, mnt := range arr {
		prt = append(prt, mnt)
	}
	return prt
}

//注意此处的坑
func test2() []*int {
	var prt []*int
	a, b, c, d, e := 1, 2, 3, 4, 5
	arr := []*int{&a, &b, &c, &d, &e}

	for _, mtn := range arr {
		var v int
		v = *mtn
		prt = append(prt, &v) //坑，
	}
	return prt
}

func main() {

	fmt.Println(test1())

	fmt.Println(test2())
}
