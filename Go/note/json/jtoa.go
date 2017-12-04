package main

import (
	"encoding/json"
	"fmt"
)

//func Unmarshal(data []byte, v interface{})  error {
//
//}

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIp":"127.0.0.1"},
	{"serverName":"Beijing_VPN","serverIp":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s.Servers) //[{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2}]

	fmt.Println("END")
}
