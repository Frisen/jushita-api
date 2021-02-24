package service

import (
	"fmt"
	"jushita-api/config"
	"log"

	"github.com/nilorg/go-opentaobao"
)

type OrderService struct {
	Res interface{}
}

func (s *OrderService) GetTraces(data map[string]interface{}) {
	opentaobao.AppKey = fmt.Sprintf("%v", config.GetValue("appkey"))
	opentaobao.AppSecret = fmt.Sprintf("%v", config.GetValue("appsecret"))
	opentaobao.Session = fmt.Sprintf("%v", config.GetValue("session"))
	opentaobao.Router = fmt.Sprintf("%v", config.GetValue("router"))
	method := fmt.Sprintf("%v", data["method"])
	r, err := opentaobao.Execute(method, opentaobao.Parameter{
		"tid": fmt.Sprintf("%v", data["tid"]),
	})
	if err != nil {
		log.Println(err.Error())
	}
	s.Res = r.Interface()
}
