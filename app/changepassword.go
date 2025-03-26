package app

import (
	"Project-Go/model" // สมมติว่า Customer โมเดลของคุณอยู่ในแพคเกจนี้
	"Project-Go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ฟังก์ชันสำหรับเปลี่ยนรหัสผ่าน
func ChangePassword(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/changepassword")
	{
		routes.POST("/", func(c *gin.Context) {
			var changeData struct {
				Email       string `json:"email" binding:"required,email"`
				OldPassword string `json:"old_password" binding:"required"`
				NewPassword string `json:"new_password" binding:"required"`
			}

			// Binding ค่า JSON ที่ส่งมา
			if err := c.ShouldBindJSON(&changeData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var customer model.Customer
			// ค้นหาผู้ใช้งานจากอีเมล์
			result := db.Where("email = ?", changeData.Email).First(&customer)
			if result.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
				return
			}

			// ตรวจสอบรหัสผ่านเก่ากับที่เก็บในฐานข้อมูล
			if !utils.CheckPassword(customer.Password, changeData.OldPassword) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Old password is incorrect"})
				return
			}

			// เข้ารหัสรหัสผ่านใหม่
			hashedNewPassword, err := utils.HashPassword(changeData.NewPassword)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
				return
			}

			// อัปเดตรหัสผ่านใหม่ในฐานข้อมูล
			customer.Password = hashedNewPassword
			db.Save(&customer)

			c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
		})
	}
}
