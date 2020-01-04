package account

import (
	"com/lovelypet/constant"
	"com/lovelypet/middleware"
	"com/lovelypet/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Sign(router *gin.Engine) {
	sign := router.Group(constant.SignPath)
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
		res,err := response.Make(constant.SUCCESS,constant.SIGNUP_SUCCESS, gin.H{
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

		token,tErr := token(c,tel)

		var code = constant.SUCCESS
		var msg  = constant.SIGNIN_SUCCESS
		var data = gin.H{}
		if tErr == nil{
			data["tel"] = tel
			data["pwd"] = pwd
			data["token"] = token
		}else {
			code = constant.FATAL
			msg = tErr.Error()
		}

		res,err := response.Make(code,msg,data)

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

func token(c *gin.Context,tel string) (string,error)  {
	claims := middleware.LovelyClaims{
		Tel:            tel,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 过期时间 一小时
			Issuer:    "lovelypet",              //签名的发行者
		},
	}
	return middleware.Jwt.GenerateToken(claims)
}

