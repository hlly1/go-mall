package main

import (
	"go-mall/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UserRouter(r)
	router.MchntRouter(r)
	r.Run(":8888")
}