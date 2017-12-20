package main

import "fmt"

//数组是一个元素类型相同的定长的序列
//数组的容量和长度是一样的。cap() 函数和 len() 函数均输出数组的容量（即长度）。
//数字有3种创建方式:
// 1）[length]Type
// 2）[N]Type{value1,value2,...,valueN}
// 3）[...]Type{value1,value2,...,valueN}
func arr()  {
	var arr1 [5]int32
	var arr2 [5]int32 = [5]int32{1,2,3,4,5}
	arr3 := [5]int32{1,2,3,4,5}
	arr4 := [5]int32{6,7,8,9,10}
	arr5 := [...]int32{11,12,13,14,15}
	arr6 := [4][4]int32{{1},{1,2},{1,2,3}}
	arr7 := [5]string{"aaa", `bb`, "可以了", "你好，go语言", "()" }

	fmt.Println(arr1," ",arr2," ",arr3," ",arr4," ",arr5," ",arr6)
	fmt.Println(len(arr4)," ",cap(arr4))
	fmt.Println(arr7)
	for i := range arr7{
		fmt.Println(i, " ", arr7[i])
	}
}

func main()  {
	arr()
}
