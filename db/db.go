/*package Mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin/models"
)

var DB *gorm.DB

func init()  {
	var err error
	sqlStr := "root:123@tcp(localhost:3306)/demo?charset=utf8&parseTime=true&loc=Local"
	DB,err = gorm.Open(mysql.Open(sqlStr),&gorm.Config{})

	if err != nil {
		panic(err)
		//fmt.Println("数据库连接错误", err)
	}
	err = DB.AutoMigrate(&models.DemoOrder{})
	if err != nil {
		panic(err)
	}
}*/
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin/models"
)

func NewDb() *gorm.DB {
	sqlStr := "root:123@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.DemoOrder{})
	if err != nil {
		panic(err)
	}

	return db
}
