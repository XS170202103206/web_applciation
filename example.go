package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

//var db *gorm.DB
/*var err error*/

type DemoOrder struct{
	gorm.Model
	OrderNo  string `json:"order_no"`
	UserName string `json:"username"`
	Amount   float64 `json:"amount"`
	Status  string `json:"status"`
	FileUrl string `json:"file_url"`
}

func main() {
	db,err:= gorm.Open("mysql","root:123@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		//panic("连接数据库失败")
		return
	}

	db.AutoMigrate(&DemoOrder{})

	r := gin.Default()
	r.GET("/",GetProjects)

	r.Run(":8080")


}
func GetProjects(c *gin.Context){
	var order []DemoOrder
	db,_:= gorm.Open("mysql","root:123@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=true&loc=Local")
	if err := db.Find(&order).Error;err!=nil{
		c.AbortWithStatus(404)
		fmt.Println(err)
	}else{
		c.JSON(200,order)
	}
}
