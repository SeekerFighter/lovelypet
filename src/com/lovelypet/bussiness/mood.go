package bussiness

import (
	"com/lovelypet/constant"
	"com/lovelypet/response"
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
*心情业务处理
 */

func Mood(router *gin.Engine) {
	mood := router.Group(constant.MoodPath)
	{
		mood.POST(constant.MoodSubmit, submit())
		mood.POST(constant.MoodDelete, _delete())
	}
}

func submit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if res,err := response.Make(constant.SUCCESS,"提交成功");err == nil {
			c.JSON(http.StatusOK,res)
		}
	}
}

func _delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
