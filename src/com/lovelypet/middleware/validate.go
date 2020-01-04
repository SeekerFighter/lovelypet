package middleware

import (
	"com/lovelypet/constant"
	"com/lovelypet/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//请求参数校验
func ParamValid()gin.HandlerFunc  {
	return func(c *gin.Context) {

		//客户端,android,ios,web等
		if channel := c.PostForm("channel");channel == ""{
			if res,err := response.Make(constant.FATAL,fmt.Sprintf(constant.PARAM_LOST,"channel"));err == nil{
				c.JSON(http.StatusOK,res)
			}
			c.Abort()
			return
		}

		//当前使用的服务器版本
		if channel := c.PostForm("apiVer");channel == ""{
			if res,err := response.Make(constant.FATAL,fmt.Sprintf(constant.PARAM_LOST,"apiVer"));err == nil{
				c.JSON(http.StatusOK,res)
			}
			c.Abort()
			return
		}
		c.Next()
	}
}