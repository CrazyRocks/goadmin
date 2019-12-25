package tools

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func SetCache(key string, value interface{}) {
	g.Redis().Do("SET", key, value)
}

func GetCache(key string) interface{} {
	result, err := g.Redis().Do("GET", key)
	if err != nil {
		return ""
	}
	return result
}

func DelCache(key string) {
	g.Redis().Do("DEL", key)
}

func Publish(key string, value interface{}) {
	data, _ := gjson.Encode(value)
	_, err := g.Redis().Do("PUBLISH", key, data)
	if err != nil {
		glog.Printf("订阅%s发送消息失败", key)
	}
}

func Subscribe(key string) interface{} {
	result, err := g.Redis().Do("Subscribe", key)
	if err != nil {
		glog.Printf("订阅发送消息失败")
	}
	return result
}
