package component

import (
	"github.com/aiqiu506/aiks/container"
	"github.com/aiqiu506/aiks/utils"
	"log"
	"net/rpc/jsonrpc"
)

type JsonRpcStruct struct {
	Host string
	Port string
}

var JsonRpc JsonRpcStruct

func init() {
	//注册组件
	container.ComponentCI.RegisterComponent("jsonRPC", &JsonRpc)
}

func (j *JsonRpcStruct) Init(config interface{}) {
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, j)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func (j *JsonRpcStruct) Request(method string, jsonData interface{}) (interface{}, error) {
	var ret interface{}
	conn, err := jsonrpc.Dial("tcp", j.Host+":"+j.Port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.Call(method, jsonData, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
