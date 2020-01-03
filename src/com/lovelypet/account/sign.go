package account

import (
	"com/lovelypet/constant"
	"com/lovelypet/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sign(router *gin.Engine) {
	sign := router.Group(constant.Dirpath)
	{
		sign.POST(constant.Signup, signUp())
		sign.POST(constant.Signin, signIn())
		sign.POST(constant.Signout, signOut())
	}
}

//账号注册
func signUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		tel := c.PostForm("tel")
		name := c.PostForm("name")
		pwd := c.PostForm("password")
		res,err := result.Response(constant.SUCCESS,constant.SIGNUP_SUCCESS, gin.H{
			"tel":      tel,
			"name":     name,
			"password": pwd,
		})

		if err == nil {
			c.JSON(http.StatusOK,res)
		}else {
			fmt.Println(err)
		}
	}
}

//账号登录
func signIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		tel := c.PostForm("tel")
		pwd := c.PostForm("password")

		res,err := result.Response(constant.SUCCESS,constant.SIGNIN_SUCCESS, gin.H{
				"tel":     tel,
				"password": pwd,
			})

		if err == nil {
			c.JSON(http.StatusOK,res)
		}else {
			fmt.Println(err)
		}
	}
}

//账号登出
func signOut() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
