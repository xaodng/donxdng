package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"net/http"
	"strings"
	"errors"
	"io/ioutil"
)

const (
	QuerySupportedVersion = "Type=QuerySupportedVersion\n\n"
	QueryVersion     		 = "Type=QueryVersion\n\n"
	QueryStaticStatus    	 = "Type=QueryStaticStatus\n\n"
	QueryRunStatus       	 = "Type=QueryRunStatus\n\n"
	DataTime              	 = "Type=DataTime\n\n"
	HeartBeat          	 = "Type=HeartBeat\n\n"
	
	Host = "http://172.16.11.90/sendmail2/mail.php"
	FROM = "neeq_gateway_monitor"
	TITLE = "NEEQ MESSAGE"
)

var (
	user string
)

func main()  {
	args := os.Args
	var server_info string

	if len(args) == 2{
		server_info = "127.0.0.1:50001"
		user = args[1]
	}else if len(args) == 3{
		server_info = args[1]
		user = args[2]
	}else {
		fmt.Println("parameter Err!\n")
		return
	}
	//conn, err := net.Dial("tcp",server_info)
	//设置建立连接超时
	conn, err := net.DialTimeout("tcp", server_info, time.Second*2)
	if err != nil {
		fmt.Println("Error dialing", err.Error())

		mes := fmt.Sprintf("from=%s&user=%s&title=%s&msg=HQGATEWAY LISTEN ERR FOR MINITOR!\n%s",FROM, user,TITLE,err.Error())
		httpPost(Host, mes)
		return
	}
	channel := make(chan  struct{})
	go sendHeartBeat(conn, channel)

	rbuf := make([]byte, 1024*100)
	var sbuf string
	times := time.Now().Format("15:04:05")
	for ;times > "08:30:00" && times < "15:05:00";times = time.Now().Format("15:04:05"){

		<-channel

		fmt.Println("system time:",times)

		_, err = conn.Write([]byte(HeartBeat + QueryStaticStatus + QueryRunStatus))
		if err != nil {
			fmt.Println("Write err!")
			continue
		}
		rn, err := conn.Read(rbuf)
		if err != nil {
			fmt.Println("Read err!")
			continue
		}

		sbuf = string(rbuf[:rn])
		fmt.Println("[----",sbuf,"]")
		if str := strHandle(server_info, sbuf);len(str) > 0 {
			fmt.Println("---",str)
			httpPost(Host,str)
		}

		//time.Sleep(time.Second * 20)
	}

		fmt.Println("End")
}

//发送心跳包
func sendHeartBeat(conn net.Conn, channel chan<- struct{})  {
	rbuf := make([]byte, 1024*100)
	var num int = 0

	for ; ;num++ {
		if num >= 3{//即6次发送查询包
			channel <- struct{}{}
			num = 0
		}

		_,err := conn.Write([]byte(HeartBeat))
		if err != nil{
			fmt.Println("Write err!")
			return
		}
		rn, err := conn.Read(rbuf)
		if err != nil {
			fmt.Println("Read err!")
			return
		}
		fmt.Println(string(rbuf[:rn]))
		time.Sleep(time.Second * 3) //每十秒发送一次
	}

}

func strHandle(server_info, data string) string {
	line := strings.Split(data, "\r\n")
	mes := make(map[string]string)

	for _,param := range line{
		fmt.Println("line----",line)
		slice1 := strings.Split(param, "=")
		fmt.Println("slice1----",slice1)

		if len(slice1) < 2 {
			continue
		}
		mes[slice1[0]]  = slice1[1]
		fmt.Println(slice1[0], " | ", slice1[1])
	}
	fmt.Println("Len = ", len(mes))
	for key, value := range mes{
		fmt.Printf("%s:%s\n",key, value)
	}

	comm1, ok1 := mes["tcpCommStatus"]
	comm2, ok2 := mes["ucpCommStatus"]

	fmt.Printf("%v, %#v \n",ok1, comm1)
	fmt.Printf("%v, %#v \n",ok2, comm2)

	if ok1 && ok2 &&(comm1 == "BAD" && comm2 == "BAD"){
		return fmt.Sprintf("from=%s&user=%s&title=%s&msg=%s HQGATEWAY CONNECTION ERR!\n%s",FROM, user, TITLE, server_info, data)
	}
	return ""
}

func httpPost(url, data string) error {
	//这里定义链接和post的数据
	//函数 strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return errors.New("http Post err!")// 将字符串 text 包装成一个 error 对象返回
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	fmt.Println(string(body))

	return nil
}