package component

import (
	"aiks/utils"
	"github.com/aiqiu506/aiks/container"
	"log"
	"net/rpc/jsonrpc"
)

type JsonRpcStruct struct {
	Host string
	Port string
}

var JsonRpc JsonRpcStruct

func init(){
	//注册组件
	container.ComponentCI.RegisterComponent("jpush", &JsonRpc)
}

func (j *JsonRpcStruct )Init(config interface{}){
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, j)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func (j * JsonRpcStruct) Request(method string, jsonData interface{}) {
	var ret bool
	conn, err := jsonrpc.Dial("tcp", j.Host+":"+j.Port)
	if err != nil {
		log.Fatalln("rpc连接错误:" + err.Error())
	}
	defer conn.Close()
	err = conn.Call(method, jsonData, &ret)
	if err != nil {
		log.Fatalln("rpc服务错误：" + err.Error())
	}
}