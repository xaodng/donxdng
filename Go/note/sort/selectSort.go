package main

import "fmt"

type SortInterface interface {
	sort()
}

type Sortor struct {
	name string
}

func main()  {
	array := []int{6, 1, 3, 5, 8, 4, 2, 0, 9, 7}
	selectsort := Sortor{name:"选择排序--从小到大--不稳定-n*n--"}
	selectsort.sort(array)

	fmt.Println(selectsort.name, array)
}

func (sorter Sortor) sort(arry []int)  {
	arrylength := len(arry)
	for i := 0; i < arrylength; i++{
		min := i
		for j := i + 1;j < arrylength; j++{
			if arry[j] < arry[i]{
				min = j
			}
		}
	arry[i], arry[min] = arry[min], arry[i]
	}
}