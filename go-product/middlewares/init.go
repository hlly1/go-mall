package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// The 1st letter of func name should Capital otherwise it would not be imported as public
func Middle_ware_test(ctx *gin.Context) {
	println("I am middle ware.")

	//This will execute Create() of UserController, then execute codes back below itself
	ctx.Next()

	//This will jump the controller to execute codes below itself
	// ctx.Abort()
	println("I am the last one to execute")
}

func IsLogin(ctx *gin.Context) {

	//Assume get session or token from redis or somewhere...
	// ctx.Set("username", "Alice")

	// _, ok := ctx.Get("username")
	session := sessions.Default(ctx)
	sessionName := session.Get("username")
	
	if sessionName == nil {
		ctx.Abort()
		ctx.JSON(200, "user not logged in")
	}

}
