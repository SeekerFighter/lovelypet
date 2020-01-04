package main

import (
	"com/lovelypet/account"
	"com/lovelypet/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lovely pet server start ...")
	router := gin.Default()
	router.Use(middleware.ParamValid())
	account.Sign(router)

	err := router.Run()
	if err != nil {
		 fmt.Println("lovely pet sever start error",err)
	}
}
