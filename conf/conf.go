package conf

import (
	"github.com/aiqiu506/aiks/utils"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type ConfigEngine struct {
	data map[interface{}]interface{}
}
// 将ymal文件中的内容进行加载
func (c *ConfigEngine) Load (path string) error {
	ext := c.guessFileType(path)
	if ext == "" {
		return errors.New("cant not load" + path + " config")
	}
	return c.loadFromYaml(path)
}

//判断配置文件名是否为yaml格式
func (c *ConfigEngine) guessFileType(path string) string {
	s := strings.Split(path,".")
	ext := s[len(s) - 1]
	switch ext {
	case "yaml","yml":
		return "yaml"
	}
	return ""
}

// 将配置yaml文件中的进行加载
func (c *ConfigEngine) loadFromYaml(path string) error {
	yamlS,readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	// yaml解析的时候c.data如果没有被初始化，会自动为你做初始化
	err := yaml.Unmarshal(yamlS, &c.data)
	if err != nil {
		return errors.New("can not parse "+ path + " config" )
	}
	return nil
}

// 从配置文件中获取值
func (c *ConfigEngine) Get(name string) interface{}{
	path := strings.Split(name,".")
	data := c.data
	for key, value := range path {
		v, ok := data[value]
		if !ok {
			break
		}
		if (key + 1) == len(path) {
			return v
		}
		if reflect.TypeOf(v).String() == "map[interface {}]interface {}"{
			data = v.(map[interface {}]interface {})
		}
	}
	return nil
}

// 从配置文件中获取string类型的值
func (c *ConfigEngine) GetString(name string) string {
	value := c.Get(name)
	switch value:=value.(type){
	case string:
		return value
	case bool,float64,int:
		return fmt.Sprint(value)
	default:
		return ""
	}
}

// 从配置文件中获取int类型的值
func (c *ConfigEngine) GetInt(name string) (int,error) {
	value := c.Get(name)
	if value==nil{
		return 0,errors.New("配置项不存在")
	}
	switch value := value.(type){
	case string:
		i,err:= strconv.Atoi(value)
		return i,err
	case int:
		return value,nil
	case bool:
		if value{
			return 1,nil
		}
		return 0,nil
	case float64:
		return int(value),nil
	default:
		return 0,errors.New("未知的配置项数据类型")
	}
}

// 从配置文件中获取bool类型的值
func (c *ConfigEngine) GetBool(name string) (bool,error) {
	value := c.Get(name)
	if value==nil{
		return false,errors.New("配置项不存在")
	}
	switch value := value.(type){
	case string:
		str,err:= strconv.ParseBool(value)
		return str,err
	case int:
		if value != 0 {
			return true,nil
		}
		return false,nil
	case bool:
		return value,nil
	case float64:
		if value != 0.0 {
			return true,nil
		}
		return false,nil
	default:
		return false,errors.New("未知的配置项数据类型")
	}
}

// 从配置文件中获取Float64类型的值
func (c *ConfigEngine) GetFloat64(name string) (float64,error) {
	value := c.Get(name)
	if value==nil{
		return 0,errors.New("配置项不存在")
	}
	switch value := value.(type){
	case string:
		str,err := strconv.ParseFloat(value,64)
		return str,err
	case int:
		return float64(value),nil
	case bool:
		if value {
			return float64(1),nil
		}
		return float64(0),nil
	case float64:
		return value,nil
	default:
		return float64(0),errors.New("未知的配置项数据类型")
	}
}

// 从配置文件中获取Struct类型的值,这里的struct是你自己定义的根据配置文件
func (c *ConfigEngine) GetStruct(name string,s interface{}) interface{}{
	d := c.Get(name)
	if d==nil{
		return errors.New("配置项不存在")
	}
	switch d.(type){
	case string:
		return  c.GetString(name)

	case map[interface{}]interface{}:
		er:=utils.MapToStruct(d.(map[interface{}]interface{}),s)
		if er!=nil{
			log.Fatalln(er)
		}
	}
	return s
}
