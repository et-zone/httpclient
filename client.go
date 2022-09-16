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
	DefaultTimeout =time.Duration(10)
)

var cli *http.Client

type client struct {
	*http.Client
	param
}
func GetClient()client{
	return client{cli,newParam()}
}

var timeout = time.Second * 10

func NewDefaultPool() {
	cli =&http.Client{
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
		Timeout: time.Second * DefaultTimeout,
	}
}
func init() {
	cli = http.DefaultClient
}

func (this *client) Dao(method string, url string, body []byte) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url can not nil")
	}
	if len(this.param.args) != 0 {
		url = url + "?"
		for _, key := range this.param.args {
			url += key + "=" + this.param.vals[key] + "&"
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
	return b, nil
}

func (this *client) Get(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url can not nil")
	}
	if len(this.param.args) != 0 {
		url = url + "?"
		for _, key := range this.param.args {
			url += key + "=" + this.param.vals[key] + "&"
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

	return b, nil
}

//
//func (this *Client) Close() {
//	this.CloseIdleConnections()
//}

type param struct {
	args      []string
	vals      map[string]string
	headerMap map[string]string
}

func newParam() param {

	return param{[]string{}, map[string]string{}, map[string]string{}}
}

func (this *param) SetParam(key string, val string) *param {
	this.args = append(this.args, key)
	this.vals[key] = val
	return this
}

func (this *param) SetHeader(key string, val string) *param {
	this.headerMap[key] = val
	return this
}

func (this *param) GetHeader() map[string]string {

	return this.headerMap
}

func (this *param) GetParam(key string) string {
	return this.vals[key]
}
