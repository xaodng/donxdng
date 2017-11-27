package main

import (
	"os"
	"net"
	"fmt"
	"time"
)

func main()  {
	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp",service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for{
		conn,err := listener.Accept()
		if err != nil{
			continue
		}
		go handleClient(conn)
		//result, err := ioutil.ReadAll(conn)
		//fmt.Println(string(result))
	}
}

func handleClient(conn net.Conn)  {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	fmt.Println(daytime)
	fmt.Println("this is server:return to client!")
}

func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}