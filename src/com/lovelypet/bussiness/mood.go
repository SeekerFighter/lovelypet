package bussiness

import (
	"com/lovelypet/constant"
	"github.com/gin-gonic/gin"
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

	}
}

func _delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
