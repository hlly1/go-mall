package main

import (
	"go-user/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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

	//session middleware
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("redis-session-secret-key"))
	r.Use(sessions.Sessions("iSession", store))

	router.UserRouter(r)
	router.MchntRouter(r)

	r.GET("/testRes", func(ctx *gin.Context) {
		// data1 := Tuple{}
		// data1.Key = "key1"
		// data1.Value = map[string]any{"content": 1111, "other": "woho!"}
		// data2 := Tuple{}
		// data2.Key = "key2"
		// data2.Value = map[string]any{"content": 2222, "other": "woho!"}
		// ctx.JSON(200, multipleArgs(gin.H{"code": 0000, "msg": "响应成功"}, data1, data2))

		test := gin.H{"code": 0000, "msg": "响应成功"}
		test["data1"] = map[string]any{"content": 1111, "other": "woho!"}
		ctx.JSON(200, test)
	})

	r.Run(":8888")
}
