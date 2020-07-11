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
	client := httpclient.InitClient("http://httpbin.org/get")
	client.Param.SetParam("name", "213").SetParam("aaa", "222")
	rep, _ := client.Get()
	fmt.Println(string(rep))
}
func Post() {
	// post
	client := httpclient.InitClient("http://httpbin.org/post")
	rep, _ := client.Dao("POST", []byte("hahidshdsfad返回"))
	fmt.Println(string(rep))
}
