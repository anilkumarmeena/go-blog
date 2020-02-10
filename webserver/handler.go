package webserver

import (
	"github.com/valyala/fasthttp"
)

func handleAPIRequest(ctx *fasthttp.RequestCtx, path Route) {
	if path == ROOT {
		ctx.Error("hii your programme is working", 400)
	}
}
