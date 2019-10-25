package component

import (
	"github.com/aiqiu506/aiks/container"
	"github.com/aiqiu506/aiks/utils"
	"github.com/go-redis/redis"
	"log"
)

type RedisParams struct {
	Host string `map:"host"`
	Port string `map:"port"`
	Auth string `map:"auth"`
	DB   int    `map:"db"`
}

type redisStruct struct {
	DB *redis.Client
}

var Redis redisStruct

func init() {
	//注册组件
	container.ComponentCI.RegisterComponent("redis", &Redis)
}
func (r *redisStruct) Init(config interface{}) {
	redisParams := &RedisParams{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, redisParams)
		if err != nil {
			log.Fatal(err)
		}
	}
	r.DB = RedisConnect(redisParams)
}

func RedisConnect(r *RedisParams) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Auth, // no password set
		DB:       r.DB,   // use default DB
	})
	if redis == nil {
		log.Fatalln("redis初始化错误")
	}
	return redis
}
