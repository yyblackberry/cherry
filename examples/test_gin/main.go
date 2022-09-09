package main

import (
	"github.com/cherry-game/cherry"
	cherryGin "github.com/cherry-game/cherry/components/gin"
)

func main() {
	testApp := cherry.NewApp("../config/", "local", "web-1")
	defer testApp.OnShutdown()

	httpServer := cherryGin.NewHttp("web_1", testApp.Address())
	httpServer.Use(cherryGin.Cors(), cherryGin.MaxConnect(2))

	httpServer.Register(new(Test1Controller))

	testApp.Startup(httpServer)
}