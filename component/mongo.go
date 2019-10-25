package component

import (
	"fmt"
	"github.com/aiqiu506/aiks/container"
	"github.com/aiqiu506/aiks/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type MongoParams struct {
	Host   string `map:"host"`
	Port   string `map:"port"`
	DBName string `map:"db"`
	User   string `map:"user"`
	Pwd    string `map:"pwd"`
}

type mongoStruct struct {
	Session *mgo.Session
	DB      *mgo.Database
}

var Mongo mongoStruct

func init() {
	//注册组件
	container.ComponentCI.RegisterComponent("mongo", &Mongo)
}

func (m *mongoStruct) Init(config interface{}) {
	mongoParams := &MongoParams{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, mongoParams)
		if err != nil {
			log.Fatal(err)
		}
	}
	m.DB, m.Session = MongoConnect(mongoParams)
}

func MongoConnect(mongoConfig *MongoParams) (*mgo.Database, *mgo.Session) {
	//mongo连接格式串  mongodb://user:pass@host1:port1
	var conStr string
	if mongoConfig.User != "" {
		conStr = fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoConfig.User, mongoConfig.Pwd, mongoConfig.Host, mongoConfig.Port)
	} else {
		conStr = fmt.Sprintf("%s:%s", mongoConfig.Host, mongoConfig.Port)
	}

	mongoSession, err := mgo.Dial(conStr)
	if err != nil {
		log.Fatalln("mongo连接错误:" + err.Error())
	}
	mongoSession.SetMode(mgo.Monotonic, true)
	db := mongoSession.DB(mongoConfig.DBName)
	if db == nil {
		log.Fatalln("mongo连接错误1:" + err.Error())
	}
	return db, mongoSession
}

func (m *mongoStruct) UseDB(dbName string) *mgo.Database {
	return m.Session.DB(dbName)
}
func (m *mongoStruct) Insert(table string, d bson.M) (bool, error) {
	err = m.DB.C(table).Insert(d)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *mongoStruct) Update(table string, selector bson.M, data bson.M) (bool, error) {
	err = m.DB.C(table).Update(selector, bson.M{"$set": data})
	if err != nil {
		return false, err
	}
	return true, nil
}
func (m *mongoStruct) IsExist(table string, selector bson.M) (bool, error) {
	cnt, err := m.DB.C(table).Find(selector).Count()
	if err != nil {
		return false, err
	}
	if cnt > 0 {
		return true, nil
	}
	return false, nil
}
