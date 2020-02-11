package webserver

import (
	"github.com/valyala/fasthttp"
)

func handleAPIRequest(ctx *fasthttp.RequestCtx, path Route) {
	if path == ROOT {
		ctx.SendFile("static/index.html")
	}
	if path == 2 {
		ctx.SendFile("static" + string(ctx.Path()))
	}
}
