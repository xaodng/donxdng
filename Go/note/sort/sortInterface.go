package main

import (
	"sort"
	"fmt"
)


// 自定义的 Reverse 类型
type Reverse struct {
	sort.Interface // 这样，Reverse可以接纳任何实现了sort.Interface的对象
}

// Reverse 只是将其中的 Inferface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func main()  {
	ints := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(ints)
	fmt.Println("after sort by Ints:\t", ints)

	doubles := []float64{2.3, 3.2, 6.7, 10.9, 5.9, 1.8}
	sort.Float64s(doubles)
	fmt.Println("after sort by Float64s:\t", doubles)

	strings := []string{"hello", "good", "students", "morning", "people", "world"}
	sort.Strings(strings)
	fmt.Println("after sort by Strings:\t", strings)

	ipos := sort.SearchInts(ints, 5)
	fmt.Printf("pos of 5 is %d th\n", ipos)

	dpos := sort.SearchFloat64s(doubles, 6.7)
	fmt.Printf("pos of 5.0 is %d th\n", dpos)

	fmt.Printf("double is asc ? %v\n", sort.Float64sAreSorted(doubles))

	doubles = []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}
	// sort.Sort(sort.Float64Slice(doubles)) // float64 排序方法 2
	// fmt.Println("after sort by Sort:\t", doubles) // [3.5 4.2 8.9 20.14 79.32 100.98]
	(sort.Float64Slice(doubles)).Sort()   // float64 排序方法 3
	fmt.Println("after sort by Sort:\t", doubles)  // [3.5 4.2 8.9 20.14 79.32 100.98]

	sort.Sort(Reverse{sort.Float64Slice(doubles)}) // float64 逆序排序
	fmt.Println("after sort by Reversed Sort:\t", doubles)




}
