package router

import (
	"go-user/middlewares"

	"go-user/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	userAuthRoutes := r.Group("/validate")
	{

		userAuthRoutes.GET("/login", user.UserController{}.Login)

	}

	// userRoutes := r.Group("/user"	)
	// The 1st way to execute middleware for this group globally
	userRoutes := r.Group("/user", middlewares.IsLogin)
	{
		// guess it will render a page to client...but frontend is not integrated this time

		// userRoutes.GET("/", func(ctx *gin.Context) {
		// 	ctx.HTML(200, "whatever", nil)
		// })

		userRoutes.GET("/create", user.UserController{}.Create)

		userRoutes.GET("/get", user.UserController{}.Index)

		userRoutes.GET("/update", user.UserController{}.Edit)

		userRoutes.GET("/logout", user.UserController{}.Logout)

	}
}
