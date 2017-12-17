package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	//"time"
	//"encoding/json"
)

func main()  {
	c, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connetc to redis error", err)
		return
	}
	//fmt.Println("Connect redis-server success")
	defer c.Close()

	_, err = c.Do("lpush", "runoobkey", "redis")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mongodb")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = c.Do("lpush", "runoobkey", "mysql")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	value,_ := redis.Values(c.Do("lrange", "runoobkey","0", "100"))
	for _,v := range value {
		fmt.Println(string(v.([]byte)))
	}

	//key := "profile"
	//imap := map[string]string{"username":"666","phonenumber":"888"}
	//value,_ := json.Marshal(imap)
	//
	//fmt.Println(string(value))
	//
	//n, err := c.Do("SETNX",key,value)

	//if err != nil {
	//	fmt.Println(err)
	//}
	//if n == int64(1) {
	//	fmt.Println("sucess")
	//}
	//
	//var imapGet map[string]string
	//
	//valueGet, err := redis.Bytes(c.Do("get", key))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//errShal := json.Unmarshal(valueGet, &imapGet)
	//if errShal != nil {
	//	fmt.Println(err)
	//}
	//
	//
	//fmt.Println(imapGet["username"])
	//fmt.Println(imapGet["phonenumber"])

	//读写
	//_, err = c.Do("SET","mykey","xaodngShi", "EX", 5)
	//_, err = c.Do("SET","mykey","xaodngShi")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}

	//username, err := redis.String(c.Do("GET","mykey"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//}else {
	//	fmt.Printf("Get mykey:%v \n", username)
	//}



	//设置过期时间
	//time.Sleep(8 * time.Second)
	//
	//username, err = redis.String(c.Do("GET","mykey"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//}else {
	//	fmt.Printf("Get mykey:%v \n", username)



	//检测值是否存在
	//is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey"))
	//if err != nil {
	//	fmt.Println("error:",err)
	//}else{
	//	fmt.Printf("exists or not: %v \n", is_key_exit)
	//}

	// 删除key
	//_, err = c.Do("DEL", "mykey")
	//if err != nil {
	//	fmt.Println("redis delete failed:", err)
	//}
	//
	//username, err = redis.String(c.Do("get", "mykey"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//}else {
	//	fmt.Printf("Get mykey:%v \n", username )
	//}

}
