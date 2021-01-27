package main

import (
	"fmt"
	"io/ioutil"

	"github.com/nilorg/go-opentaobao"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func init() {
	var c conf
	con := c.getconf()
	log.Println("appkey------>>>>>>", con.Session)
	opentaobao.AppKey = con.AppKey
	opentaobao.AppSecret = con.AppSecret
	opentaobao.Session = con.Session
	opentaobao.Router = con.Router
}

func main() {
	method := "taobao.jds.trade.traces.get"
	res, err := opentaobao.Execute(method, opentaobao.Parameter{
		"tid": "1234139580066577",
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("订单轨迹", res)
}

type conf struct {
	AppKey    string
	AppSecret string
	Session   string
	Router    string
}

func (c *conf) getconf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
