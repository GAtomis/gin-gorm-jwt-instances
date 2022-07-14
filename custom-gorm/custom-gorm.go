/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-11 16:03:01
 * @LastEditTime: 2022-07-14 17:08:11
 * @LastEditors: Gavin
 */
package customgorm

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
)

//定义一个数据库类型结构
type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100);"`
}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers"`
}
type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}
type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers"`
}

//JTW模版
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//创建一个令牌 jwt
func initJWT() (string, error) {
	mySigningKey := []byte("")

	c := MyClaims{
		Username: "Gavin",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //开始时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //过期时间
			Issuer:    "Gavin",                     //戳
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := token.SignedString(mySigningKey) //加密
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return "", err
	} else {
		fmt.Println(s)
		return s, nil
	}

}

//解析一个令牌
func parseToken(tokenString string) {
	mySigningKey := []byte("")
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

}
func initSQL(fn func(*gorm.DB)) {

	db, _ := gorm.Open("mysql", "root:root@tcp(ec2-3-112-56-234.ap-northeast-1.compute.amazonaws.com:3306)/artemis?charset=utf8mb4&parseTime=True&loc=Local")
	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})
	// i := IDCard{
	// 	Num: 123456,
	// }
	// s := Student{
	// 	StudentName: "zhounan",
	// 	IDCard:      i,
	// }
	// t := Teacher{
	// 	TeacherName: "老师傅",
	// 	Students: []Student{
	// 		s,
	// 	},
	// }
	// c := Class{
	// 	ClassName: "迈瑞中国",
	// 	Students: []Student{
	// 		s,
	// 	},
	// }
	// _ = db.Create(&c).Error
	// _ = db.Create(&t).Error
	fn(db)
	defer db.Close()
}

//模拟中间件2
func middleTow() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before next2")
		c.Next()
		fmt.Println("after next2")
	}
}
func studentByPost(s *Student) {
	initSQL(func(db *gorm.DB) {
		db.Create(s)
	})
}
func studentByGet(id string) Student {
	var student Student
	initSQL(func(db *gorm.DB) {

		db.Preload("Teachers").Preload("IDCard").First(&student, "id=?", id)

	})
	return student
}
func classByGet(id string) Class {
	var class Class
	initSQL(func(db *gorm.DB) {

		db.First(&class, "id=?", id)

	})
	return class
}
func classByPost(c *Class) {

	initSQL(func(db *gorm.DB) {
		db.Create(c)

	})

}
func OnServer() {
	r := gin.Default()
	testModel := r.Group("test").Use(middleTow())
	testModel.POST("student", func(ctx *gin.Context) {
		var s Student
		err := ctx.ShouldBindJSON(&s)

		if err != nil {
			ctx.JSON(200, gin.H{
				"msg":  "报错了",
				"data": []any{},
			})
		} else {
			studentByPost(&s)
			ctx.JSON(200, gin.H{
				"msg":  "success",
				"data": s,
			})
		}

	})
	testModel.GET("student", func(ctx *gin.Context) {

		id := ctx.Query("id")

		if id == "" {
			ctx.JSON(200, gin.H{
				"msg":  "报错了",
				"data": []any{},
			})
		} else {
			s := studentByGet(id)
			ctx.JSON(200, gin.H{
				"msg":  "success",
				"data": s,
			})
		}

	})
	testModel.GET("class", func(ctx *gin.Context) {

		id := ctx.Query("id")

		if id == "" {
			ctx.JSON(200, gin.H{
				"msg":  "报错了",
				"data": []any{},
			})
		} else {
			s := classByGet(id)
			ctx.JSON(200, gin.H{
				"msg":  "success",
				"data": s,
			})
		}

	})
	testModel.POST("class", func(ctx *gin.Context) {
		var c Class
		err := ctx.ShouldBindJSON(&c)

		if err != nil {
			ctx.JSON(200, gin.H{
				"msg":  "报错了",
				"data": []any{},
			})
		} else {
			fmt.Printf("c: %v\n", c)
			classByPost(&c)
			ctx.JSON(200, gin.H{
				"msg": "success",
			})
		}

	})

	testModel.GET("token", func(ctx *gin.Context) {
		s := ctx.GetHeader("Authorization")

		parseToken(s)
		ctx.JSON(200, gin.H{
			"msg":  "报错了",
			"data": s,
		})
	})
	r.Run(":9527")

}

func (u User) TableName() string {
	return "am_user"
}
