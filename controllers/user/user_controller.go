package user

import "github.com/gin-gonic/gin"

type UserController struct {
}

// (con UserController) means mount the method to the struct
func (con UserController) Create(ctx *gin.Context) {
	println("ctx.Next() executes me now...") 
	ctx.JSON(200, "create succeed!")
}

func (con UserController) Index(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	name, _ := username.(string)
	ctx.JSON(200, "[succeed] username is "+name)
}

func (con UserController) Edit(ctx *gin.Context) {
	ctx.JSON(200, "edit succeed!")
}
