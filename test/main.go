package main

import (
	"fmt"

	"github.com/et-zone/httpclient"
)

func main() {
	Post()
	//Get()
}
func Get() {
	client:=httpclient.GetClient()
	//get

	client.SetParam("name", "213").SetParam("aaa", "222")

	for i := 0; i < 1; i++ {
		rep, _ := client.Get("http://www.baidu.com")
		fmt.Println(string(rep))
	}

}
func Post() {
	// post
	client := httpclient.GetClient()
	client.SetHeader("fff", "ggg")
	for i := 0; i < 1; i++ {
		client.Dao( "POST", "http://www.baidu.com", []byte("hahidshdsfad返回"))

	}
}
