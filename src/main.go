package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lovely pet server start ...")
	router := gin.Default()

	err := router.Run()
	if err != nil {
		 fmt.Println("lovely pet sever start error",err)
	}
}
