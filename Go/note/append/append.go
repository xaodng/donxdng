package main

import "fmt"

func test1() []int  {
	var prt []int
	arr := []int{1, 2, 3, 4, 5}

	for _,mnt := range arr{
		prt = append(prt,mnt)
	}
	return prt
}

func test2() []*int  {
	var prt []*int
	a,b,c,d,e := 1, 2, 3, 4, 5
	arr := []*int{&a, &b, &c, &d, &e}

	for _,mnt := range arr{
		v := *mnt
		fmt.Print(v, " ") //1 2 3 4 5
		fmt.Print(mnt, " ") //0x123821f0 0x123821f4 0x123821f8 0x123821fc 0x12382200
		//prt = append(prt,mnt)
		prt = append(prt,&v) //此处注意，有坑
	}
	fmt.Println()
	return prt
}

func main()  {

	fmt.Println(test1()) //[1 2 3 4 5]

	fmt.Println(test2()) //[0x12382204 0x12382208 0x1238220c 0x12382220 0x12382224]
}