package main

import "fmt"

// channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或 多个goroutine之间传递消息。

func main()  {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 1)
	go sum(values[:len(values)/2],resultChan)
	go sum(values[len(values)/2:],resultChan)
	go sum(values[len(values)/3:],resultChan)

	sum1, sum2, sum3 := resultChan, resultChan, resultChan
	fmt.Println("Result:",sum1, sum2, sum3)

}

func sum(values []int, resultChan chan int)  {
	sum := 0;
	for _,value := range values {
		sum += value
	}
	resultChan <-sum
}