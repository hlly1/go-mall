package router

import "github.com/gin-gonic/gin"

func MchntRouter(r *gin.Engine) {
	mchntRoute := r.Group("/mchnt")
	{
		mchntRoute.GET("/getInfo", func(ctx *gin.Context) {
			ctx.JSON(200, "Mc'Donald - I'm Lovin it")
		})
	}
}
