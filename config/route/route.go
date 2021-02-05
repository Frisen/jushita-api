package route

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()
	order := r.Group("/order")
	{
		order.GET("/his", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	return r
}
