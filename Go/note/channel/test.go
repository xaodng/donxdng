package main

var a string

/*channel是go语言中的同步工具，有两种模式
缓冲:程序执行序列将阻塞在读channel的调用处 <- chan； 或当channel满时，阻塞在写channel调用处 chan <-。*/
var c = make(chan int, 10)

/*非缓冲:程序执行序列将阻塞在读和写channel的调用处 chan <- 或 <- chan*/
//var c = make(chan int)

func f()  {
	a = "hello,world"
	c <- 0
	//<- c
}

func main()  {
	go f()
	<- c
	//c <- 0
	print(a)
}