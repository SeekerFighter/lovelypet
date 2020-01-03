package main

import (
	"com/lovelypet/account"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lovely pet server start ...")
	router := gin.Default()

	account.Sign(router)

	err := router.Run()
	if err != nil {
		 fmt.Println("lovely pet sever start error",err)
	}
}
