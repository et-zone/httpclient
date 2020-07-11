package httpclient

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	http.Client
	Url string
	Param
	ctime  time.Time
	method string
}

var timeout = time.Second * 10

func InitClient(baseUrl string) *Client {
	client := &Client{Client: http.Client{}}
	client.Timeout = timeout
	client.Url = baseUrl
	client.Param = Param{args: []string{}, vals: map[string]string{}}
	client.ctime = time.Now()
	return client
}
func (this *Client) Dao(method string, body []byte) ([]byte, error) {
	this.method = method
	url := this.Url
	if len(this.Param.args) != 0 {
		url = url + "?"
		for _, key := range this.Param.args {
			url += key + "=" + this.Param.vals[key] + "&"
		}
		url = url[:len(url)-1]
	}
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	res, err := this.Do(req)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	log.Println("Method:", this.method, "Path: ", this.Url, " time: ", time.Since(this.ctime))
	return b, nil
}

func (this *Client) Get() ([]byte, error) {
	this.method = "GET"
	url := this.Url
	if len(this.Param.args) != 0 {
		url = url + "?"
		for _, key := range this.Param.args {
			url += key + "=" + this.Param.vals[key] + "&"
		}
		url = url[:len(url)-1]
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	res, err := this.Do(req)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	log.Println("Method:", this.method, "Path: ", this.Url, " time: ", time.Since(this.ctime))
	return b, nil
}

type Param struct {
	args []string
	vals map[string]string
}

func (this *Param) SetParam(key string, val string) *Param {
	this.args = append(this.args, key)
	this.vals[key] = val
	return this
}

func (this *Param) GetParam(key string) string {
	return this.vals[key]
}
