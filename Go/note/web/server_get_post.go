package main

import (
	"net/http"
	"fmt"
	"html"
	//"strings"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Server struct {
	ServerName	string
	ServerIP 	string
}

type Serverslice struct {
	Servers []Server
	ServersID	string
}

func main()  {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hi,I love you %s",html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET"{
		fmt.Println("method:", r.Method)
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

	for k, v := range r.Form {
		fmt.Print("key:", k, "; ")
		fmt.Println("val:", strings.Join(v, ""))
		}
	}else if r.Method == "POST" {
		result,_ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)

		//未知类型的处理办法
		var f interface{}
		json.Unmarshal(result, &f)

		m := f.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "if int", vv)
			case float64:
				fmt.Println(k, "if flar64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv{
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		//结构已知， 解析到结构体
		var s Serverslice
		json.Unmarshal([]byte(result), &s)

		fmt.Println(s.ServersID)
	}
}
