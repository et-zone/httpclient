package main

import (
	"fmt"

	"github.com/et-zone/httpclient"
)

func main() {
	//Post()
	Get()
}
func Get() {
	//get
	client := httpclient.InitDefaultClient()
	client.Param.SetParam("name", "213").SetParam("aaa", "222")
	client.AppName = "ggg"
	for i := 0; i < 1; i++ {
		rep, _ := client.Get(httpclient.NewContext(), "http://www.baidu.com")
		fmt.Println(string(rep))
	}

}
func Post() {
	// post
	client := httpclient.InitDefaultClient()
	client.Param.SetHeader("fff", "ggg")
	for i := 0; i < 1; i++ {
		client.Dao( "POST", "http://www.baidu.com:8888/ping", []byte("hahidshdsfad返回"))

	}
}
