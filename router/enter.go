/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-08-04 14:37:42
 * @LastEditTime: 2022-08-04 15:21:37
 * @LastEditors: Gavin
 */
package router

import (
	"encoding/json"
	"fmt"
	"gin-web/common"
	"gin-web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const httpport = "9999"

func InitRouter() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	rg := router.Group("api")
	rg.POST("getToken", GetToken)
	router.Run(":" + httpport)

}

func GetToken(c *gin.Context) {

	fmt.Println("111")
	tokenReq := &model.GetTokenReq{}
	err := c.BindJSON(tokenReq)
	if err != nil {
		c.JSON(http.StatusOK, &model.Resp{Code: 5001, Msg: err.Error()})
		return
	}

	rtmtoken, err := common.BuildToken(tokenReq.Appid, tokenReq.AppSecert,
		tokenReq.Userid, 1, 0)

	tokenrsp := model.Resp{Code: 0, Msg: "ok", Ts: 1, RequestId: "123343", Data: rtmtoken}

	str, _ := json.Marshal(tokenrsp)

	fmt.Println("****libs.SendGift(*sendGiftReq)**** ", string(str))

	c.JSON(http.StatusOK, tokenrsp)
}
