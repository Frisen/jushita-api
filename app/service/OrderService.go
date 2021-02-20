package service

import (
	"jushita-api/config"
	"log"

	"github.com/bitly/go-simplejson"
	"github.com/nilorg/go-opentaobao"
)

type OrderService struct {
	Res *simplejson.Json
}

func (s *OrderService) GetTraces(data map[string]interface{}) {
	opentaobao.AppKey = config.GetValue("appkey").(string)
	opentaobao.AppSecret = config.GetValue("appsecret").(string)
	opentaobao.Session = config.GetValue("session").(string)
	opentaobao.Router = config.GetValue("router").(string)
	method := data["method"].(string)
	r, err := opentaobao.Execute(method, opentaobao.Parameter{
		"tid": data["tid"].(string),
	})
	if err != nil {
		log.Println(err.Error())
	}
	s.Res = r
}
