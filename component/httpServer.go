package component

import (
	"aiks/component/AiksHttpServer"
	"github.com/aiqiu506/aiks/container"
	"github.com/aiqiu506/aiks/utils"
	"log"
)

type HTTPServerParams struct {
	Port string `map:"port"`
	Host string `map:"host"`
}

type httpServer struct {
	Server *AiksHttpServer.Server
}

func (h * httpServer) Init(config interface{}) {
   params:=&HTTPServerParams{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf,params)
		if err != nil {
			log.Fatal(err)
		}
	}
	h.Server=h.ServerRun(params)

}

func (h * httpServer)ServerRun(p *HTTPServerParams) *AiksHttpServer.Server{
	return AiksHttpServer.New(p.Host,p.Port)
}

var HttpServer httpServer

func init(){
	//注册日志组件
	container.ComponentCI.RegisterComponent("httpServer", &HttpServer)
}
