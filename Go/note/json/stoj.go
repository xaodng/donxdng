package main

import (
	"encoding/json"
	"os"
	//"fmt"
)

type Server struct {
	ID int `json:"id,string"`
	ServerName string `json:"serverName"`
	ServerName2 string `json:"serverName2"`
	ServerIP string `json:"serverIP"`
}
/*
针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:
字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的
时候会把该字段对应的值转换成JSON字符串
 */
func main()  {
	s := Server{
		ID: 3,
		ServerName: `Go "1.0"`,
		ServerName2:`Go "1.0"`,
		ServerIP:``,
	}

	b,_ := json.Marshal(s)
	//fmt.Println(string(b))
	os.Stdout.Write(b)
}