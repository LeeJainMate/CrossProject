package app

import (
	"Project-Go/model" // สมมติว่า Customer โมเดลของคุณอยู่ในแพคเกจนี้
	"Project-Go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชัน Login ที่รับการตรวจสอบข้อมูลผู้ใช้งาน
func Login(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/login")
	{
		routes.POST("/", func(c *gin.Context) {
			var loginData struct {
				Email    string `json:"email" binding:"required,email"`
				Password string `json:"password" binding:"required"`
			}

			if err := c.ShouldBindJSON(&loginData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var customer model.Customer
			result := db.Where("email = ?", loginData.Email).First(&customer)

			if result.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
				return
			}

			// ตรวจสอบรหัสผ่านโดยใช้ bcrypt
			if !utils.CheckPassword(customer.Password, loginData.Password) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
				return
			}

			// หากรหัสผ่านถูกต้อง ส่งข้อมูลทั้งหมดกลับ
			c.JSON(http.StatusOK, gin.H{
				"user": gin.H{
					"id":         customer.CustomerID,
					"first_name": customer.FirstName,
					"last_name":  customer.LastName,
					"email":      customer.Email,
					"phone":      customer.PhoneNumber,
					"address":    customer.Address,
					"created_at": customer.CreatedAt,
					"updated_at": customer.UpdatedAt,
				},
			})
		})
	}
}
