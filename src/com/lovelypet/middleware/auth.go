package middleware

import (
	"com/lovelypet/constant"
	"com/lovelypet/response"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	TokenExpired      = errors.New("Token is expired")
	TokenNotValidYet  = errors.New("Token not active yet")
	TokenMalformed    = errors.New("That's not even a token")
	TokenInvalid      = errors.New("Couldn't handle this token:")
	Jwt *JWT
)

func init() {
	Jwt = newJWT("lovelyToken")
	fmt.Println("auth init()called,JWT init success...")
}

type LovelyClaims struct {
	Tel   string `json:"tel"`
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

func newJWT(key string) *JWT {
	return &JWT{
		[]byte(key),
	}
}

func (j *JWT)GenerateToken(claims LovelyClaims)(string,error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT)ParseToken(tokenStr string) (*LovelyClaims,error)  {
	token,err := jwt.ParseWithClaims(tokenStr,&LovelyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey,nil
	})
	if err != nil {
		if ve,ok := err.(*jwt.ValidationError);ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims,ok := token.Claims.(*LovelyClaims);ok&&token.Valid {
		return claims,nil
	}
	return nil,TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &LovelyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*LovelyClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.GenerateToken(*claims)
	}
	return "", TokenInvalid
}

func AccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("lovelyToken")
		if token == "" {
			res,err := response.Make(constant.FATAL,constant.AUTH_TOKEN_LOST)
			if err == nil {
				c.JSON(http.StatusOK,res)
			}
			c.Abort()
			return
		}

		claims,tErr := Jwt.ParseToken(token)
		if tErr != nil {
			if res,err := response.Make(constant.FATAL,tErr.Error());err == nil {
				c.JSON(http.StatusOK,res)
			}
			c.Abort()
			return
		}
		c.Set("claims",claims)
		c.Next()
	}
}