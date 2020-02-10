package main

import (
	"code.slike.in/golang/go-blog/webserver"
	"code.slike.in/golang/go-server-boilerplate/app"
)

func main() {
	server := app.CreateFastHttp(webserver.HandleRequest)
	application := app.CreateAppWithServer(server)
	application.Start()
}
