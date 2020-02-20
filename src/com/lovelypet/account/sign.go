package account

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/constant"
	"lovelypet/src/com/lovelypet/middleware"
	"lovelypet/src/com/lovelypet/model"
	"lovelypet/src/com/lovelypet/response"
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
		fmt.Println("Signup()called:name=",name,",tel=",tel,",pwd=",pwd)
		var(
			res gin.H
			err error
		)
		u := model.NewUser(name,tel,pwd)
		if u.IsUserExist() {
			res,err = response.Make(constant.FATAL,constant.SignupRepeat)
		}else if u.Insert() {
			res,err = response.Make(constant.SUCCESS,constant.SignupSuccess, gin.H{
				"tel":      tel,
				"name":     name,
				"password": pwd,
			})
		}else {
			res,err = response.Make(constant.FATAL,constant.ServerSqlError)
		}
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
		fmt.Println("signIn()called: tel=",tel,"pwd=",pwd)
		var code = constant.SUCCESS
		var msg  = constant.SigninSuccess
		var data = gin.H{}
		u := model.NewUser("",tel,"")

		if u.IsUserExist() {
			if u.Pwd == pwd {
				token,tErr := token(c,tel)
				if tErr == nil{
					data["tel"] = tel
					data["pwd"] = pwd
					data["token"] = token
				}else {
					code = constant.FATAL
					msg = tErr.Error()
				}
			}else {
				code = constant.FATAL
				msg = constant.PwdError
			}
		}else {
			code = constant.FATAL
			msg = constant.SignupNotRecord
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