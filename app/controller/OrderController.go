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

func (o *OrderController) GetOrderHis(c *gin.Context) {
	o.display = &common.Display{Context: c}
	o.data = common.GetData(c)
	o.os = new(service.OrderService)
	defer o.display.CatchPanic()
	switch {
	case o.data["method"] == "taobao.jds.trade.traces.get":
		if o.data["tid"] != nil {
			o.get()
		}
	}
}

func (o *OrderController) get() {
	o.os.GetTraces(o.data)
	o.display.Show(o.os.Res)
}
