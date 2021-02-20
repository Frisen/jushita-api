package controller

import (
	"jushita-api/app/common"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	display *common.Display
	data    map[string]interface{}
}

func (s *IndexController) Run(c *gin.Context) {
	s.display = &common.Display{Context: c}
	s.display.Show("welcome here bro.")
}
