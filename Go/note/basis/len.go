package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	s := "Go编程"

	fmt.Println("---------unsafe---------")
	fmt.Println(unsafe.Sizeof(true)) //bool 1 Byte
	fmt.Println(unsafe.Sizeof(int8(0))) // 1 Byte
	fmt.Println(unsafe.Sizeof(int16(10))) // 2 Byte
	fmt.Println(unsafe.Sizeof(int(1000000000))) // 4 Byte
	fmt.Println(unsafe.Sizeof(int32(10000000))) // 4 Byte
	fmt.Println(unsafe.Sizeof(rune(1000000000))) // 4 Byte
	fmt.Println(unsafe.Sizeof(int64(10000000000000))) // 8 Byte
	fmt.Println(unsafe.Sizeof(s)) // 8 Byte
	fmt.Println("---------unsafe---------")

	fmt.Println(len(s)) //8 Byte,每个汉字占3个Byte

	fmt.Println(len([]rune(s))) // 4个

	fmt.Println(len(string(rune('编')))) //每个汉字 3 Byte

	fmt.Println(len(string(rune('A')))) //1 Byte

	for _,c := range s{ // 注意for ... range string, 返回的c是rune/int32类型
		fmt.Printf("%T, %c\n",c, c)
	}
}
