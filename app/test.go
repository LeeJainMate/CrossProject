package app

import "github.com/gin-gonic/gin"

func DemoController(router *gin.Engine) {
	routes := router.Group("/ping")
	{
		routes.GET("/", ping)

	}

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong!!! Pong2",
	})
}
