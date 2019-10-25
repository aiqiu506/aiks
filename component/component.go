package component

import (
	"aiks/container"
	"log"
)

var (

)

func Init(componentName string, config interface{})  {
	if component,ok:=container.ComponentCI.Handel[componentName];ok{
		component.Init(config)
	}else{
		log.Fatal("组件【"+componentName+"】不存在")
	}

}