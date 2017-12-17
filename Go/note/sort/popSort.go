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
	array := []int{6, 1, 3, 5, 8, 4, 2, 0, 9, 7}
	popsort := Sortor{name:"冒泡排序--从小到大--不稳定-n*n--"}
	popsort.sort(array)

	fmt.Println(popsort.name, array)
}

func (sorter Sortor) sort(arry []int)  {
	arrylength := len(arry)
	for i := 0; i < arrylength; i++{
		for j := i + 1; j < arrylength; j++{
			if arry[i] > arry[j] {
				arry[i] ,arry[j] = arry[j], arry[i]
			}
		}
	}
}