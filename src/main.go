package main

import (
	. "fmt"
	"github.com/gin-gonic/gin"
	"lovelypet/src/com/lovelypet/account"
	"lovelypet/src/com/lovelypet/bussiness"
	"lovelypet/src/com/lovelypet/cache"
	"lovelypet/src/com/lovelypet/middleware"
	"lovelypet/src/com/lovelypet/model"
)

func main() {

	Println("lovely pet server start ...")

	cache.CreateTable(&model.User{})

	router := gin.Default()
	router.Use(middleware.HeaderSet())
	router.Use(middleware.ParamValid())
	account.Sign(router)
	router.Use(middleware.AccessToken())
	{
		bussiness.Mood(router)
	}

	err := router.Run()

	if err != nil {
		 Println("lovely pet sever start error",err)
	}
}
