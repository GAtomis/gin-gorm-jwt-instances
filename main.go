/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-06 22:37:08
 * @LastEditTime: 2022-08-04 14:43:42
 * @LastEditors: Gavin
 */
package main

import (
	"fmt"

	"gin-web/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//json body  uri /:id/:name  form query  binding:表单验证
type Person struct {
	Name string `json:"name" uri:"name" form:"name" binding:"required"`
	Age  int    `json:"age" uri:"age" form:"age" binding:"required,big"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex"`
}

type Pet struct {
	gorm.Model
	Type string
	Sex  bool
	Age  int
	Name string
}

//参数验证函数
func big(f1 validator.FieldLevel) bool {
	fmt.Printf("f1: %T\n", f1.Field().Interface().(int))
	if f1.Field().Interface().(int) <= 18 {
		println("true")
		return false
	}
	return true

}

//模拟中间件 Next函数前为请求前  1=>2=>core=>2=>1
func middleOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		age := c.Query("age")
		fmt.Printf("age: %T\n", age)
		if len(age) != 0 {
			c.Next()
		} else {
			c.Abort()
		}

		fmt.Println("after next")
	}
}

//模拟中间件2
func middleTow() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("before next2")
		c.Next()
		fmt.Println("after next2")
	}
}
func main() {

	// r := gin.Default()

	// customgorm.OnServer()
	router.InitRouter()

	// db, _ := gorm.Open("mysql", "root:yfqdmr@tcp(ec2-18-183-20-186.ap-northeast-1.compute.amazonaws.com:3306)/artemis?charset=utf8mb4&parseTime=True&loc=Local")
	// db.AutoMigrate(&Pet{})
	//更新 单一更新
	// db.Model(&Pet{}).Where("id=?", 1).Update("name", "张筱楠")
	//创建
	// db.Model(&Pet{}).Create(&Pet{
	// 	Name: "周楠",
	// 	Sex:  false,
	// 	Type: "忧郁蓝猫",
	// 	Age:  30,
	// })
	//软删除
	// db.Where("id in (?)", []int{1, 2}).Delete(&Pet{})
	//逻辑删除
	// db.Where("id in (?)", []int{1, 2}).Unscoped().Delete(&Pet{})

	// defer db.Close()

	// fmt.Println("数据库成功")
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("big", big)
	// }
	// r.GET("/good/:id", func(c *gin.Context) {
	// 	username := c.DefaultQuery("username", "zhounan")
	// 	id := c.Param("id")
	// 	password := c.Query("password")

	// 	c.JSON(200, gin.H{"username": username, "id": id, "password": password})

	// })
	// r.POST("/id", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "pong "})
	// })
	// r.PUT("/good/:id", func(ctx *gin.Context) {
	// 	s := ctx.Param("id")
	// 	ctx.JSON(200, gin.H{"message": s})
	// })
	// //JSON body
	// r.POST("/bind", func(ctx *gin.Context) {
	// 	var p Person
	// 	err := ctx.ShouldBindJSON(&p)
	// 	if err != nil {
	// 		ctx.JSON(200, gin.H{
	// 			"msg": "报错了",
	// 		})
	// 	} else {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": p,
	// 		})

	// 	}
	// })
	// //uri 传参数
	// r.POST("/uri/:name/:age/:sex", func(ctx *gin.Context) {
	// 	var p Person
	// 	err := ctx.ShouldBindUri(&p)
	// 	if err != nil {
	// 		ctx.JSON(200, gin.H{
	// 			"msg": "报错了",
	// 		})
	// 	} else {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": p,
	// 		})

	// 	}
	// })
	// //query入参
	// r.POST("/query", func(ctx *gin.Context) {
	// 	var p Person
	// 	err := ctx.ShouldBindQuery(&p)
	// 	if err != nil {
	// 		fmt.Printf("错误日志:%v", err.Error())
	// 		ctx.JSON(200, gin.H{
	// 			"msg": "报错了",
	// 		})
	// 	} else {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": p,
	// 		})

	// 	}
	// })
	// //上传文件
	// r.POST("/uploadFile", func(ctx *gin.Context) {
	// 	fh, _ := ctx.FormFile("file")
	// 	ctx.SaveUploadedFile(fh, "./"+fh.Filename)
	// 	ctx.JSON(200, gin.H{
	// 		"msg":  "上传成功",
	// 		"file": fh,
	// 	})

	// })
	// //分组路由
	// v1 := r.Group("v1").Use(middleOne(), middleTow())
	// v1.GET("test", func(ctx *gin.Context) {
	// 	var p Person
	// 	err := ctx.ShouldBindQuery(&p)
	// 	if err != nil {
	// 		fmt.Printf("错误日志:%v", err.Error())
	// 		ctx.JSON(200, gin.H{
	// 			"msg": "报错了",
	// 		})
	// 	} else {
	// 		ctx.JSON(200, gin.H{
	// 			"msg":  "success",
	// 			"data": p,
	// 		})

	// 	}
	// })

	// r.Run(":9527")

}
