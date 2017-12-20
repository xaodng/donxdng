package main

import "fmt"

//切片是长度可变、容量固定的相同的元素序列
//切片本质是一个数组。容量固定是因为数组的长度是固定的，切片的容量即隐藏数组的长度。
// 长度可变指的是在数组长度的范围内可变。
//切片有4种创建方式:
//1）make([]Type, length, capacity)
//2）make([]Type, length)
//3）[]Type{}
//4）[]Type{value1, value2,...,valueN}
//创建切片跟创建数组唯一的区别在于 Type 前的“ [] ”中是否有数字，为空，则代表切片，否则则代表数组。
// 因为切片是长度可变的。
func slice()  {
	s1 := make([]int32, 5, 8)
	s2 := make([]int32, 9)
	s3 := []int32{}
	s4 := []int32{1,2, 3, 4, 5}
	fmt.Println(s1," ",s2," ",s3," ",s4)

	myslice := make([]int, 5, 10)
	myslice = append(myslice, 1, 2, 3)
	myslice2 := [] int{8, 9, 10, 11, 12}
	myslice = append(myslice, myslice2...)
	fmt.Println(myslice)

}

//切片与隐藏数组
//一个切片是一个隐藏数组的引用，并且对于该切片的切片也引用同一个数组。
func slice1()  {
	s0 := []string{"a","b","c","d","e"}
	s1 := s0[2:len(s0)-1] //[c d]
	s2 := s0[:3] //[a b c]

	fmt.Println(s0, s1, s2) //[a b c d e] [c d] [a b c]
	s2[2] = "8"
	fmt.Println(s0, s1, s2) //[a b 8 d e] [8 d] [a b 8]
	//切片s0 、 s1和 s2是同一个底层数组的引用，所以s2改变了，其他两个都会变
}

//遍历、修改切片
//只有用索引才能修改元素。如在第一个遍历中，赋值ele为7，结果没有作用。
// 因为在元素遍历中，ele是值传递，ele是该切片元素的副本，修改它不会影响原本值，
// 而在第四个遍历——索引遍历中，修改的是该切片元素引用的值，所以可以修改。
func slice2()  {
	s0 := []string{"a","b","c","d","e"}
	fmt.Println("------元素遍历-------")
	for _,ele := range s0 {
		fmt.Print(ele, " ")
		ele = "7"
	}
	fmt.Println(s0)
	fmt.Println("\n------索引遍历------")
	for index := range s0 {
		fmt.Print(index,":",s0[index], " ")
	}
	fmt.Println("\n------元素索引共用------")
	for index, ele := range s0 {
		fmt.Print(index,":",ele, s0[index], " ")
	}
	fmt.Println("\n------修改-------")
	for index := range s0 {
		s0[index] = "9"
	}
	fmt.Println(s0)
}
//追加、复制切片
//追加、复制切片，用的是内置函数append和copy，copy函数返回的是最后所复制的元素的数量
func slice3()  {
	slice := []int32{}
	fmt.Printf("len=%d,slice:%v\n",len(slice),slice)

	slice = append(slice, 12,11,10,9)
	fmt.Printf("追加后,len=%d,slice:%v\n",len(slice),slice)

	slicecp := make([]int32, (len(slice)))
	n := copy(slicecp,slice)
	fmt.Printf("复制后,len=%d,slice:%v,%d\n",len(slicecp),slicecp,n)
}

//内置函数append
//内置函数append可以向一个切片后追加一个或多个同类型的其他值。
// 如果追加的元素数量超过了原切片容量，那么最后返回的是一个全新数组中的全新切片。
// 如果没有超过，那么最后返回的是原数组中的全新切片。无论如何，append对原切片无任何影响。
func slice4()  {
	slice := [6]int32{1,2,3,4,5,6}
	slice2 := slice[:2] //[1 2]
	_=append(slice2,30,40,50,60,70)
	fmt.Println(slice) //[1 2 3 4 5 6]
	fmt.Println(slice2) //[1 2]
	_=append(slice2,30,40)
	fmt.Println(slice) //[1 2 30 40 5 6]
	fmt.Println(slice2) //[1 2]
}

func main()  {
	slice()
	//slice1()
	//slice2()
	//slice3()
	//slice4()
}
