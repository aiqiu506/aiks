package component

import (
	"aiks/container"
	"aiks/utils"
	"errors"
	"github.com/DeanThompson/jpush-api-go-client"
	"github.com/DeanThompson/jpush-api-go-client/push"
	"log"
)

type JpushParams struct {
	AppKey    string `map:"appKey"`
	IsTestEnv bool   `map:"isTestEnv"`
	Secret    string `map:"secret"`
}

type jpushStruct struct {
	Client *jpush.JPushClient
	Env  string
}

var JPush jpushStruct

func init() {
	//注册组件
	container.ComponentCI.RegisterComponent("jpush", &JPush)
}

func (j *jpushStruct) Init(config interface{}) {
	jpushParams := &JpushParams{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, jpushParams)
		if err != nil {
			log.Fatal(err)
		}
	}
	j.Client = CreateClient(jpushParams)

}

func CreateClient(config *JpushParams) *jpush.JPushClient {
	return jpush.NewJPushClient(config.AppKey, config.Secret)
}

type JPushData struct {
	To       string `json:"to"`
	Platform string `json:"platform"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Extras  map[string]interface{} `json:"extras"`
}

func setPlatform(platform *push.Platform, plat string) error {
	if plat == "all" {
		err := platform.Add("ios", "android")
		return err
	}
	return nil
}

func (j *jpushStruct) Push(data JPushData) (bool, error) {

	platform := push.NewPlatform()
	//设置平台
	err := setPlatform(platform, data.Platform)

	if err != nil {
		return false, err
	}

	//设置推送给
	audience := push.NewAudience()
	audience.SetAlias([]string{data.To})

	//audience.SetAlias([]string{msg.To})

	notification := push.NewNotification(data.Name)
	androidNotification := push.NewAndroidNotification(data.Content)
	androidNotification.Title = data.Title
	androidNotification.AddExtra("data", data.Extras)

	notification.Android = androidNotification

	//iOS 平台专有的 notification，用 alert 属性初始化
	iosNotification := push.NewIosNotification(data.Content)
	iosNotification.Badge = 1
	iosNotification.AddExtra("data", data.Extras)
	iosNotification.Sound = "default"
	//Validate 方法可以验证 iOS notification 是否合法
	//一般情况下，开发者不需要直接调用此方法，这个方法会在构造 PushObject 时自动调用
	notification.Ios = iosNotification
	/*m := push.NewMessage(msg.Content)
	m.Title = msg.Title*/
	//m.Extras=
	// option 对象，表示推送可选项
	options := push.NewOptions()
	// iOS 平台，是否推送生产环境，false 表示开发环境；如果不指定，就是生产环境
	options.ApnsProduction = false
	if j.Env != "production" {
		options.ApnsProduction = true
	}

	// Options 的 Validate 方法会对 time_to_live 属性做范围限制，以满足 JPush 的规范
	//	options.TimeToLive = 10000000
	// Options 的 Validate 方法会对 big_push_duration 属性做范围限制，以满足 JPush 的规范
	//	options.BigPushDuration = 1500

	payload := push.NewPushObject()
	payload.Platform = platform
	payload.Audience = audience
	payload.Notification = notification
	/*payload.Message = m*/
	payload.Options = options

	// Push 会推送到客户端
	result, err := j.Client.Push(payload)
	//	showErrOrResult("Push", result, err)

	// PushValidate 的参数和 Push 完全一致
	// 区别在于，PushValidate 只会验证推送调用成功，不会向用户发送任何消息
	//result, err := client.PushValidate(payload)
	if err != nil {
		return false,err
	}
	if result.StatusCode != 200 {

		return false,errors.New(result.Error.Message)
	}
	return true,nil
}
