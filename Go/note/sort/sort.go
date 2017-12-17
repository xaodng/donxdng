package main

import (
	"sort"
	"fmt"
)

func main()  {
	intList := []int{3, 5, 1, 7, 8, 4, 6, 2, 0, 9}
	float8List := []float64{4.2, 8.4, 3.5, 94.7}
	stringList := []string{"t", "y", "u", "i", "o", "j", "k", "l", "z", "x", "v", "c"}

	//升序排序
	//对于 int 、 float64 和 string 数组或是分片的排序，
	// go 分别提供了 sort.Ints() 、 sort.Float64s() 和 sort.Strings() 函数， 默认都是从小到大排序。
	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n", intList)
	fmt.Printf("%v\n", float8List)
	fmt.Printf("%v\n", stringList)

	//逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Printf("%v\n", intList)
	fmt.Printf("%v\n", float8List)
	fmt.Printf("%v\n", stringList)

}
