package app

// import (
// 	"Project-Go/model"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	// สมมติว่ามีโมเดล User ที่ใช้กับ GORM หรือฐานข้อมูล
// 	"gorm.io/gorm"
// )

// func Login(router *gin.Engine, db *gorm.DB) {
// 	routes := router.Group("/login")
// 	{
// 		routes.POST("/", func(c *gin.Context) {
// 			// รับข้อมูล email และ password จาก body ของ request
// 			var loginData struct {
// 				Email    string `json:"email" binding:"required,email"`
// 				Password string `json:"password" binding:"required"`
// 			}

// 			// ตรวจสอบข้อมูล JSON ที่รับมา
// 			if err := c.ShouldBindJSON(&loginData); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			// ค้นหาผู้ใช้งานจากฐานข้อมูลโดยใช้ email
// 			var user model.Customer
// 			result := db.Where("email = ? AND password = ?", loginData.Email, loginData.Password).First(&user)

// 			// หากไม่พบผู้ใช้หรือรหัสผ่านไม่ตรง
// 			if result.Error != nil {
// 				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
// 				return
// 			}

// 			// หากพบผู้ใช้และรหัสผ่านตรง, ส่งข้อมูลทั้งหมดของผู้ใช้กลับ
// 			c.JSON(http.StatusOK, gin.H{
// 				"user": gin.H{
// 					"first_name": user.FirstName,
// 					"last_name":  user.LastName,
// 					"email":      user.Email,
// 					"phone":      user.PhoneNumber,
// 					"address":    user.Address,
// 					"created_at": user.CreatedAt,
// 					"updated_at": user.UpdatedAt,
// 				},
// 			})
// 		})
// 	}
// }

// // ฟังก์ชัน Login ที่รับการตรวจสอบข้อมูลผู้ใช้งาน
// // func Login(router *gin.Engine) {
// // 	routes := router.Group("/login")
// // 	{

// // 		routes.POST("/", pingPost)
// // 	}
// // }

// // func pingPost(c *gin.Context) {
// // 	email := c.PostForm("email")
// // 	password := c.PostForm("password")
// // 	c.JSON(201, gin.H{
// // 		"message": "Hello World!!!" + email + " " + password,
// // 	})
// // }
