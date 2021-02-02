package httpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"

	"net"
	"net/http"
	"time"
)

const (
	MaxIdleConnsPerHost = 100
	IdleConnTimeout     = 300
	MaxIdleConns        = 1000
	MaxConnsPerHost     = 1000
)

type Client struct {
	http.Client
	Param
	ctime   time.Time
	AppName string
}

var timeout = time.Second * 10

func InitDefaultClientPool() *Client {
	client := &Client{Client: http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
			Proxy:             http.ProxyFromEnvironment,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxConnsPerHost:     MaxConnsPerHost,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},

		Timeout: time.Second * 10,
	}}

	client.Param = Param{args: []string{}, vals: map[string]string{}, headerMap: map[string]string{}}
	client.ctime = time.Now()
	return client
}

func InitDefaultClient() *Client {
	client := &Client{Client: http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
			Proxy:             http.ProxyFromEnvironment,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxConnsPerHost:     1,
			MaxIdleConnsPerHost: 1,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},

		Timeout: time.Second * 10,
	}}

	client.Param = Param{args: []string{}, vals: map[string]string{}, headerMap: map[string]string{}}
	client.ctime = time.Now()
	return client
}

func InitClient(c http.Client) *Client {
	client := &Client{Client: c}
	client.Param = Param{args: []string{}, vals: map[string]string{}, headerMap: map[string]string{}}
	client.ctime = time.Now()
	return client
}

func (this *Client) Dao(ctx *eContext, method string, url string, body []byte) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url can not nil")
	}
	this.ctime = time.Now()
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

	header := this.GetHeader()
	for k, v := range header {
		req.Header.Set(k, v)
	}

	res, err := this.Do(req)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	seteContext(ctx, this.ctime, this.AppName, method, url, time.Since(this.ctime), res.StatusCode)
	logINFO(ctx)
	return b, nil
}

func (this *Client) Get(ctx *eContext, url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url can not nil")
	}
	this.ctime = time.Now()
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
	header := this.GetHeader()
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := this.Do(req)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}

	seteContext(ctx, this.ctime, this.AppName, "GET", url, time.Since(this.ctime), res.StatusCode)
	logINFO(ctx)
	return b, nil
}

func (this *Client) Close() {
	this.CloseIdleConnections()
}

type Param struct {
	args      []string
	vals      map[string]string
	headerMap map[string]string
}

func NewParam() Param {

	return Param{[]string{}, map[string]string{}, map[string]string{}}
}

func (this *Param) SetParam(key string, val string) *Param {
	this.args = append(this.args, key)
	this.vals[key] = val
	return this
}

func (this *Param) SetHeader(key string, val string) *Param {
	this.headerMap[key] = val
	return this
}

func (this *Param) GetHeader() map[string]string {

	return this.headerMap
}

func (this *Param) GetParam(key string) string {
	return this.vals[key]
}

func init() {
	err := initLog()
	if err != nil {
		panic(err.Error())
	}
}
