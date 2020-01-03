package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("lovely pet server start ...")
	router := gin.Default()
	
	router.Run()
}
