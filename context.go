package httpclient

import (
	"context"
	"strings"
	"time"
)

type eContext struct {
	context.Context
	appName  string
	method   string
	ip       string
	path     string //不需要ip地址
	nowtime  *time.Time
	duration time.Duration
}

func NewContext() *eContext {
	return &eContext{context.TODO(), "", "", "", "", nil, 0}
}

func seteContext(ctx *eContext, t time.Time, appName string, method string, path string, duration time.Duration) {
	if ctx == nil {
		return
	}
	plist := strings.Split(path, "?")

	if strings.Contains(plist[0], "http") {
		list := strings.SplitN(plist[0], "/", 4)
		ctx.ip = list[2]
		ctx.path = "/" + list[3]
	} else {
		list := strings.SplitN(plist[0], "/", 2)
		ctx.ip = list[0]
		ctx.path = "/" + list[1]
	}
	ctx.appName = appName
	ctx.method = method
	ctx.nowtime = &t
	ctx.duration = duration
}

func (ctx *eContext) GeteContextInfo() (nowtime *time.Time, appName string, ip string, mothod string, path string, duration time.Duration) {
	if ctx == nil {
		t := time.Now()
		return &t, "", "", "", "", 0
	}
	return ctx.nowtime, ctx.appName, ctx.ip, ctx.method, ctx.path, ctx.duration
}
