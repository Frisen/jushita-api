package controller

import (
	"jushita-api/app/common"
	"jushita-api/app/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	display *common.Display
	data    map[string]interface{}
	os      *service.OrderService
}

func GetOrderHis(c *gin.Context) {
	o := &OrderController{
		display: &common.Display{Context: c},
		data:    common.GetData(c),
		os:      new(service.OrderService),
	}
	defer o.display.CatchPanic()

	switch {
	case o.data["method"] == "taobao.jds.trade.traces.get":
		if o.data["tid"] != nil {
			o.get()
		}
	}
	o.display.Finish()
}

func (o *OrderController) get() {
	o.os.GetTraces(o.data)
	// fmt.Println(*o.os.Res)
	o.display.Show(*o.os.Res)
}
