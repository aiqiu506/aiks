package aiks

import (
	"aiks/component"
	"aiks/conf"
	"aiks/container"
	"log"
)

type App struct {
	configPath string
}

func NewApp() *App{
	return &App{
	}
}

var (
	CommonConf * conf.ConfigEngine
	err error
)
func (a *App)SetConfigPath(path string){
	a.configPath=path
}
//解析配置文件
func (a *App)parseConfig() *conf.ConfigEngine{
	var config conf.ConfigEngine
	err=config.Load(a.configPath)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
//初始化组件
func (a *App)InitComponent(components ... string)  {
	CommonConf=a.parseConfig()
	for _,componentName:= range components{
		config:=CommonConf.Get(componentName)
		if config!=nil{
			component.Init(componentName,config)
		}else{
			log.Fatal("组件["+componentName+"]配置不存在")
		}
	}
}

func(a *App)Start(){
	//初始化服务容器
	container.InitContainer()
	//服务执行
	container.Container.Run()
}


