package main

import (
	"go-mall/router"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	//load web files when integrated with front-end(html with vars rendered by go), should be set ahead of routers
	// r.LoadHTMLGlob("web/**/*")

	//set static web pages(html...), the 1st argument is router while the 2nd one is file path
	// r.Static("/static_pages", "./static")

	//This will execute the middle ware globaly
	// r.Use(middle_ware_test)

	router.UserRouter(r)
	router.MchntRouter(r)
	r.Run(":8888")
}