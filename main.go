package main

import (
	"fmt"

	"github.com/nilorg/go-opentaobao"
	log "github.com/sirupsen/logrus"
)

func init() {
	opentaobao.AppKey = "28293853"
	opentaobao.AppSecret = "2e6d041383be62f877968635999369f4"
	opentaobao.Session = "6100e2829efceeb2ba879f5201331fd02d711dada40c31a217101303"
	opentaobao.Router = "http://gw.api.taobao.com/router/rest"
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

	// urlVaule := url.Values{}
	// urlVaule.Add("sign", "FBA1B6E438B299A65CF0712B20FBB6CF")
	// urlVaule.Add("timestamp", "2021-01-25 18:30:00")
	// urlVaule.Add("v", "2.0")
	// urlVaule.Add("method", "taobao.jds.trade.traces.get")
	// urlVaule.Add("sign_method", "hmac")
	// urlVaule.Add("tid", "1234139580066577")
	// urlVaule.Add("force_sensitive_param_fuzzy", "true")
	// urlVaule.Add("format", "json")
	// resp, _ := http.PostForm("http://gw.api.taobao.com/router/rest", urlVaule)
	// body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
}
