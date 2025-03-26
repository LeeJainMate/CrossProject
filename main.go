package main

import (
	"Project-Go/app"
	"Project-Go/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config") // ชื่อไฟล์ config
	viper.AddConfigPath(".")      // ระบุพาธของไฟล์ config
	err := viper.ReadInConfig()   // อ่านไฟล์ config
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// ดึงค่า DSN จากไฟล์ config.yaml
	dsn := viper.GetString("mysql.dsn")

	// เชื่อมต่อกับ MySQL โดยใช้ GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ทดสอบการเชื่อมต่อ
	fmt.Println("Database connected successfully!")

	// ตรวจสอบการเชื่อมต่อ
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	customers := []model.Customer{}
	db.Find(&customers)
	// fmt.Println(customers)

	app.StartServer()
}
