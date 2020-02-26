package bussiness

import (
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/constant"
	"lovelypet/src/com/lovelypet/middleware"
	"lovelypet/src/com/lovelypet/model"
	"lovelypet/src/com/lovelypet/response"
	"lovelypet/src/com/lovelypet/util"
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

		var (
			res gin.H
			err error
		)

		userId := middleware.GetUserIdFromCookie(c)
		moodText := c.PostForm("moodText")

		moodInfo := model.NewMoodInfo(userId,moodText)

		if moodInfo.Insert() {
			if res,err = response.Make(constant.SUCCESS,moodInfo);err == nil {
				c.JSON(http.StatusOK,res)
			}
		}else {
			if res,err = response.Make(constant.FATAL,constant.ServerSqlError);err == nil {
				c.JSON(http.StatusOK,res)
			}
		}
	}
}

func _delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			res gin.H
			err error
		)

		moodId := c.PostForm("moodId")

		moodInfo := model.NewMoodInfo(0,"")

		moodInfo.ID = util.StrToUInt(moodId)

		if  moodInfo.IsMoodExist(){
			if model.Delete(moodId) {
				if res,err = response.Make(constant.SUCCESS,constant.DeleteSuccess);err == nil {
					c.JSON(http.StatusOK,res)
				}
			}else {
				if res,err = response.Make(constant.FATAL,constant.ServerSqlError);err == nil {
					c.JSON(http.StatusOK,res)
				}
			}
		}else {
			if res,err = response.Make(constant.FATAL,constant.NotRecord);err == nil {
				c.JSON(http.StatusOK,res)
			}
		}
	}
}

func _query()gin.HandlerFunc  {
	return func(c *gin.Context) {
		var (
			res gin.H
			err error
		)
		userId := middleware.GetUserIdFromCookie(c)
		moods := model.Query(userId)

		if moods  == nil {
			if res,err = response.Make(constant.FATAL,constant.ServerSqlError);err == nil {
				c.JSON(http.StatusOK,res)
			}
		}else {
			if res,err = response.Make(constant.SUCCESS,constant.QuerySuccess,moods);err == nil {
				c.JSON(http.StatusOK,res)
			}
		}
	}
}
