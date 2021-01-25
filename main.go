package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func main() {
	urlVaule := url.Values{}
	urlVaule.Add("sign", "FBA1B6E438B299A65CF0712B20FBB6CF")
	urlVaule.Add("timestamp", "2021-01-25 18:30:00")
	urlVaule.Add("v", "2.0")
	urlVaule.Add("app_key", "28293853")
	urlVaule.Add("method", "taobao.jds.trade.traces.get")
	urlVaule.Add("sign_method", "hmac")
	urlVaule.Add("session", "6100e2829efceeb2ba879f5201331fd02d711dada40c31a217101303")
	urlVaule.Add("tid", "1234139580066577")
	urlVaule.Add("force_sensitive_param_fuzzy", "true")
	urlVaule.Add("format", "json")

	resp, _ := http.PostForm("http://gw.api.taobao.com/router/rest", urlVaule)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
