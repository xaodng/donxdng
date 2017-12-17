package main

import (
	"fmt"
)

type SortInterface interface {
	sort()
}

type Sortor struct {
	name string
}

func main()  {
	array := []int{8 ,6, 1, 3, 5, 4, 2, 0, 9, 7}
	quicksort := Sortor{name:"快速排序--从小到大--不稳定-nlog2n--"}
	quicksort.sort(array)

	fmt.Println(quicksort.name, array)
}

func (sorter Sortor) sort(arry []int)  {
	if len(arry) <= 1{
		return
	}

	mid := arry[0]
	i := 1
	head, tail := 0, len(arry)-1
	for head < tail{
		if arry[i] > mid {
			arry[i], arry[tail] = arry[tail], arry[i]
			tail--
		}else {
			arry[i], arry[head] = arry[head], arry[i]
			head++
			i++
		}
	}

	arry[head] = mid
	sorter.sort(arry[:head])
	sorter.sort(arry[head+1:])


}