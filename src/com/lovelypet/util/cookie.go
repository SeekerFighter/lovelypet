package util

import (
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/constant"
)

func SetNewCookie(c *gin.Context,token string)  {
	c.SetCookie(constant.Token, token, 36000, "/", "localhost", false, true)
}

func GetToken(c *gin.Context) string {
	var token string
	token = c.Request.Header.Get(constant.Token)
	if token == "" {
		switch c.Request.Method {
		case "POST":
			token = c.PostForm(constant.Token)
		case "GET":
			token = c.Query(constant.Token)
		default:
			token = ""
		}
		if token == "" {
			token,_ = c.Cookie(constant.Token)
		}
	}
	return token
}