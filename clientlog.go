package httpclient

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"time"
)

const (
	logpath = "log/"
	logfile = "eclient.log"
)

var file *os.File
var err error

func initLog() error {
	_, err := ioutil.ReadDir(logpath)
	if err != nil {
		log.Println(err.Error())
		os.Mkdir(logpath, os.ModePerm)
	}

	file, err = os.OpenFile(logpath+time.Now().Format("2006_01_02")+"_"+logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("log_file_err", err.Error())
		return err
	}
	return nil
}

func logINFO(ctx *eContext) {
	file.WriteString("{\"time\":\"" + ctx.nowtime.Format("2006-01-02 15:04:05") +
		"\",\"appName\":\"" + ctx.appName + "\",Method\":\"" + ctx.method +
		"\",\"Path\":\"" + ctx.path + "\",\"time\":" + strconv.Itoa(int(ctx.duration)) +
		"\"code\":" + strconv.Itoa(ctx.Code) +
		"}\n")
}
