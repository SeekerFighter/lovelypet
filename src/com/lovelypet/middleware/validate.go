package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/constant"
	"lovelypet/src/com/lovelypet/response"
	"net/http"
)

//请求参数校验
func ParamValid()gin.HandlerFunc  {
	return func(c *gin.Context) {

		//客户端,android,ios,web等
		if channel := c.PostForm("channel");channel == ""{
			if res,err := response.Make(constant.FATAL,fmt.Sprintf(constant.ParamLost,"channel"));err == nil{
				c.JSON(http.StatusOK,res)
			}
			c.Abort()
			return
		}

		//当前使用的服务器版本
		if channel := c.PostForm("apiVer");channel == ""{
			if res,err := response.Make(constant.FATAL,fmt.Sprintf(constant.ParamLost,"apiVer"));err == nil{
				c.JSON(http.StatusOK,res)
			}
			c.Abort()
			return
		}
		c.Next()
	}
}

//头部设置中间件，允许跨域请求
func HeaderSet()gin.HandlerFunc  {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}