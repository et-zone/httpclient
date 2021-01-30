package main

import (
	"fmt"

	"github.com/et-zone/httpclient"
)

func main() {
	Get()
}
func Get() {
	//get
	client := httpclient.InitDefaultClient()
	client.Param.SetParam("name", "213").SetParam("aaa", "222")
	client.AppName = "ggg"
	for i := 0; i < 10; i++ {
		rep, _ := client.Get(httpclient.NewContext(), "http://127.0.0.1:8888/pong")
		fmt.Println(string(rep))
	}

}
func Post() {
	// post
	client := httpclient.InitDefaultClient()
	rep, _ := client.Dao(httpclient.NewContext(), "POST", "http://127.0.0.1:8888/ping", []byte("hahidshdsfad返回"))
	fmt.Println(string(rep))
}
