package main

import "fmt"

func main()  {
	x := min(1, 3, 2, 5) //方法一
	fmt.Printf("The minnum is %d\n", x)
	slice1 := []int{7, 9, 4, 5, 2}
	x = min(slice1...) //方法二
	fmt.Printf("The minnum in the slice is: %d\n", x)
}

//传递变长参数
func min(a ...int) int {
	fmt.Printf("type = %T, value = %#v\n",a,a)
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _,v := range a {
		if v < min{
			min = v
		}
	}
	return min
}
