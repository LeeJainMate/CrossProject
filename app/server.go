package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// เรียกใช้ฟังก์ชันที่เกี่ยวกับ Controller
	DemoController(router)
	Login(router, db) // ส่ง db ที่เชื่อมต่อแล้วไปที่ฟังก์ชัน Login
	ChangePassword(router, db)
	SearchByDescription(router, db)
	SearchByPrice(router, db)

	// เริ่มต้นเซิร์ฟเวอร์
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
