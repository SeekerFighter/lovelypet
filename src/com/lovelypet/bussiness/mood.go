package bussiness

import (
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/constant"
	"lovelypet/src/com/lovelypet/response"
	"net/http"
)
/**
*心情业务处理
 */

func Mood(router *gin.Engine) {
	mood := router.Group(constant.MoodPath)
	{
		mood.POST(constant.MoodSubmit, _submit())
		mood.POST(constant.MoodDelete, _delete())
		mood.POST(constant.MoodQuery, _query())
	}
}

func _submit() gin.HandlerFunc {
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

func _query()gin.HandlerFunc  {
	return func(c *gin.Context) {

	}
}
