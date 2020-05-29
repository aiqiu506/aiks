package aiks

import (
	"aiks/component"
	"aiks/component/AiksHttpServer"
	"github.com/aiqiu506/aiks/container"
	"sync"
	"testing"
)

type TT struct{}

func (t *TT)Run(locker *sync.WaitGroup, name string){
	defer func() {
		locker.Done()
	}()
	component.HttpServer.Server.GET("/", func(c *AiksHttpServer.Context) {
		c.JSON(200,map[string]interface{}{
			"你好":"hello",
				"世界": "world",
		})
	})
	component.HttpServer.Server.Run()
}
func TestDebugToFile(t *testing.T) {


	container.Container.RegisterService("TT", &TT{})



	app := NewApp()
	app.SetConfigPath("config.yaml")
	app.InitComponent("httpServer")
	app.Start()
}
