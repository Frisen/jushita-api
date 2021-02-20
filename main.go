package main

import (
	"jushita-api/config/route"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	(&route.Route{Engine: engine}).Run()
	engine.Run()
	// method := "taobao.jds.trade.traces.get"
	// res, err := opentaobao.Execute(method, opentaobao.Parameter{
	// 	"tid": "1234139580066577",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("订单轨迹", res)
}
