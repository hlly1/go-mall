package router

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine){
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/getInfo", func(ctx *gin.Context) {
			ctx.JSON(200, "I am Bob")
		})

	}
}