package webserver

import (
	"bytes"
	"runtime/debug"

	"code.slike.in/golang/dataservicedemo/utils"
	"github.com/kataras/golog"
	"github.com/valyala/fasthttp"
)

type Route int

const (
	NONE Route = iota
	ROOT
)

var (
	root = []byte("/")
)

func isPathValid(ctx *fasthttp.RequestCtx) (bool, Route) {
	if utils.PathIs(root, ctx) {
		return true, ROOT
	}
	return false, NONE
}

func HandleRequest(ctx *fasthttp.RequestCtx) {
	defer func() {
		if r := recover(); r != nil {
			golog.Error("error in route ")
			golog.Errorf("%s", debug.Stack())
			utils.Send200OkResponse(ctx, "error in route")
		}
	}()
	if bytes.Equal(ctx.Request.Header.Method(), []byte("OPTIONS")) {
		ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		return
	}
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

	if ok, path := isPathValid(ctx); ok {
		if path <= 2 {
			handleAPIRequest(ctx, path)
		}
	} else {
		ctx.Error("Not Found", 404)
	}
}
