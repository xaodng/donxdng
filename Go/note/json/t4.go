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