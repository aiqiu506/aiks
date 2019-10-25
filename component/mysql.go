package component

import (
	"fmt"
	"github.com/aiqiu506/aiks/container"
	"github.com/aiqiu506/aiks/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type MysqlParms struct {
	Host        string `map:"host"`
	Port        string `map:"port"`
	DBName      string `map:"db"`
	User        string `map:"user"`
	Pwd         string `map:"pwd"`
	MaxCons     int    `map:"maxCons"`
	MaxFreeCons int    `map:"MaxFreeCons"`
}

type mySqlStruct struct {
	DB *gorm.DB
}

var Mysql mySqlStruct

func init() {
	//注册组件
	container.ComponentCI.RegisterComponent("mysql", &Mysql)
}
func (my *mySqlStruct) Init(config interface{}) {
	//组件初始化
	mysqlParams := &MysqlParms{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, mysqlParams)
		if err != nil {
			log.Fatal(err)
		}
	}
	my.DB = MysqlConnect(mysqlParams)
}

func MysqlConnect(my *MysqlParms) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", my.User, my.Pwd, my.Host, my.Port, my.DBName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	if my.MaxFreeCons != 0 {
		db.DB().SetMaxOpenConns(my.MaxCons)
	}
	if my.MaxCons != 0 {
		db.DB().SetMaxIdleConns(my.MaxFreeCons)
	}
	return db
}
