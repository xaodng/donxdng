package main

import "C"
import (
	"time"
	"fmt"
)

func Seek(i int)  {
	C.srandom(C.uint(i))
}

func Rondom() int {
	return int(C.random())
}

func main()  {
	Seek(int(time.Now().Unix())) //生成随机种子

	ret := Rondom()
	fmt.Println(ret)
}