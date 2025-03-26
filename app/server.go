package app

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	DemoController(router)
	// Login(router, &gorm.DB{})
	router.Run()
}
