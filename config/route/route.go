package route

import (
	"jushita-api/app/controller"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Engine *gin.Engine
}

func (r *Route) Run() {
	r.index()
	r.order()
}

func (r *Route) index() {
	r.Engine.Any("", new(controller.IndexController).Run)
}

func (r *Route) order() {
	order := r.Engine.Group("order")
	{
		order.GET("trace", new(controller.OrderController).GetOrderHis)
	}
}
