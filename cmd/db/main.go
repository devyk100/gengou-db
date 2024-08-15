package main

import (
	"fmt"
	"github.com/devyk100/gengou-db/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	dsn := "postgresql://yash:EugslGggG1nqiGzGj9N2nA@gengou-connect-7548.6xw.aws-ap-southeast-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	//result := db.AutoMigrate(&database.User{})
	db.Create(&database.User{
		UserID:   "yash3faesfk11",
		Name:     "somet3faewhi2ng",
		UserType: database.Instructor,
		EmailID:  "yashk@32d3ev2",
		Phone:    "5156136523116251",
	})
	var user = database.User{
		UserID: "dfaw49",
	}

	db.First(&user)
	log.Println(user.Name)
	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)
	fmt.Println(now)
}
