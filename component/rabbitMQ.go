package component

import (
	"aiks/container"
	"aiks/utils"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type rabbitMQParams struct {
	Host   string `map:"host"`
	Port   string `map:"port"`
	DefaultExchange string `map:"defaultExchange"`
	User   string `map:"user"`
	Pwd    string `map:"pwd"`
}

type rabbitMQStruct struct {
	Channel      *amqp.Channel
}

var RabbitMQ rabbitMQStruct
var RabbitSession chan *amqp.Error
func init() {
	//注册组件
	container.ComponentCI.RegisterComponent("rabbitMQ", &RabbitMQ)
}

func (m *rabbitMQStruct) Init(config interface{}) {
	rabbitMQParams := &rabbitMQParams{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, rabbitMQParams)
		if err != nil {
			log.Fatal(err)
		}
	}
	m.Channel = RabbitMQConnect(rabbitMQParams)
}

func RabbitMQConnect(mq *rabbitMQParams) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", mq.User, mq.Pwd, mq.Host, mq.Port))
	if err != nil {
		log.Fatalln("rabbitMQ连接错误1："+err.Error())
	}

	mqConn, err := conn.Channel()
	if err != nil {
		log.Fatalln("rabbitMQ连接错误2："+err.Error())

	}
	if mq.DefaultExchange!=""{
		err = mqConn.ExchangeDeclare(mq.DefaultExchange, amqp.ExchangeDirect, true, false, false, false, nil)
		if nil != err {
			log.Fatalln("rabbitMQ连接错误3："+err.Error())
		}
	}
	RabbitSession=make(chan *amqp.Error)
	conn.NotifyClose(RabbitSession)
	return mqConn
}


func (mq * rabbitMQStruct)SetExchange(exchangeName string,exchangeType string)error{
	if exchangeType==""{
		exchangeType=amqp.ExchangeDirect
	}else{
		pos,err:=utils.InArray(exchangeType,[]string{amqp.ExchangeDirect,amqp.ExchangeFanout,amqp.ExchangeHeaders,amqp.ExchangeTopic})
		if err!=nil{
			log.Fatalln("exchangeType设置错误："+err.Error())
		}
		if pos==-1{
			log.Fatalln("exchangeType设置错误：应在[direct，fanout，topic，headers]其中一个")
		}
	}


	return mq.Channel.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
}

func (mq * rabbitMQStruct )Publish(data string,exchange string,queueName string) (bool,error) {
	err := mq.Channel.Publish(exchange, queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(data),
	})
	if err != nil {
		return false,err
	}
	return true,nil
}

func (mq * rabbitMQStruct )SetQueue(name string) error{

	_,err:=mq.Channel.QueueDeclare(name,true, //durable
		false,                               //delete when unused
		false,                               //exclusive
		false,                               //no wait
		nil, )
	return err

}

func (mq * rabbitMQStruct )BindQueueToExchange(queue,key,exchangeName string) error{

	err:=mq.Channel.QueueBind(
		queue,       // queue name
		key,       // routing key
		exchangeName, // exchange
		false,           // no-wait
		nil,
	)
	return err

}
func (mq * rabbitMQStruct ) Consume(queue string, callBack func(message string))error{

	msgs,err:=mq.Channel.Consume(
		queue, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err!=nil{
		return err
	}
	for msg := range msgs {
		callBack(string(msg.Body))
	}
	return nil
}
