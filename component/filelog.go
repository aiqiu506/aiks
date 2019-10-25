package component

import (
	"aiks/container"
	"aiks/utils"
	"log"
	"os"
	"time"
)

type LogParams struct {
	Path string `map:"path"`
	DefaultName string `map:"defaultName"`
	IsDaily	bool `map:"isDaily"`
	NeedDir	bool `map:"needDir"`
}

type logFile struct {
	filename string
	fileFd  *os.File
}
var Log logFile
var err error
var logParams * LogParams
func init(){
	//注册日志组件
	container.ComponentCI.RegisterComponent("log", &Log)
}

func (l *logFile) Init (config interface{}){
	logParams=&LogParams{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, logParams)
		if err != nil {
			log.Fatal(err)
		}
	}
	var name string
	if logParams.DefaultName!=""{
		name=logParams.DefaultName
	}else{
		name="log"
	}
	l.filename=logParams.makeFileName(name)
}


func (l *LogParams)makeFileName(name string) string{
	fileName:=l.Path
	//每天生成
	if l.IsDaily{
		if l.NeedDir{
			fileName+="/"+time.Now().Format("20060102")+"/"
		}else{
			name="/"+time.Now().Format("20060102")+"_"+name
		}
	}else{
		fileName+="/"
	}
	return fileName+name+".log"
}

func (l logFile) logWrite(content string,isExit bool) {
	l.fileFd=utils.OpenFile(l.filename)
	defer l.fileFd.Close()
	logFile := log.New(l.fileFd, "", log.LstdFlags)
	if isExit{
		logFile.Fatal(content)
	}else{
		logFile.Println(content)
	}
}

//调试输出
func (l logFile)OutPut(content ... interface{}){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utils.OutPutString(content...))
}
//重新修改日志文件名
func (l *logFile)SetFileName(name string) *logFile{
	 log:=*l
	 log.filename=logParams.makeFileName(name)
	return &log
}
//日志记录到文件
func (l logFile)WriteFile(content...interface{}){

	l.logWrite(utils.OutPutString(content...),false)
}
func (l logFile)WriteFileExit(content ... interface{}){
	l.logWrite(utils.OutPutString(content...),true)
}



