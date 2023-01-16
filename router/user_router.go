package router

import (
	"go-mall/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		// guess it will render a page to client...but frontend is not integrated this time
		
		// userRoutes.GET("/", func(ctx *gin.Context) {
		// 	ctx.HTML(200, "whatever", nil)
		// })

		userRoutes.GET("/create", user.UserController{}.Create)

		userRoutes.GET("/get", user.UserController{}.Index)

		userRoutes.GET("/update", user.UserController{}.Edit)

	}
}
