package app

// import (
// 	"Project-Go/model"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// func ChangePassword(router *gin.Engine, db *gorm.DB) {
// 	routes := router.Group("/change-password")
// 	{
// 		routes.POST("/", func(c *gin.Context) {
// 			var requestData struct {
// 				Email       string `json:"email" binding:"required,email"`
// 				OldPassword string `json:"old_password" binding:"required"`
// 				NewPassword string `json:"new_password" binding:"required,min=6"`
// 			}

// 			// 1️⃣ ตรวจสอบว่า Request ถูกต้องหรือไม่
// 			if err := c.ShouldBindJSON(&requestData); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			// 2️⃣ ค้นหาผู้ใช้ในฐานข้อมูล
// 			var customer model.Customer
// 			result := db.Where("email = ?", requestData.Email).First(&customer)

// 			if result.Error != nil {
// 				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
// 				return
// 			}

// 			// 3️⃣ ตรวจสอบรหัสผ่านเก่า
// 			if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(requestData.OldPassword)); err != nil {
// 				c.JSON(http.StatusUnauthorized, gin.H{"message": "Old password is incorrect"})
// 				return
// 			}

// 			// 4️⃣ เข้ารหัสรหัสผ่านใหม่
// 			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.NewPassword), bcrypt.DefaultCost)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password"})
// 				return
// 			}

// 			// 5️⃣ อัปเดตรหัสผ่านใหม่ในฐานข้อมูล
// 			customer.Password = string(hashedPassword)
// 			db.Save(&customer)

// 			// 6️⃣ ส่งข้อความแจ้งเตือนว่ารหัสผ่านเปลี่ยนสำเร็จ
// 			c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
// 		})
// 	}
// }
