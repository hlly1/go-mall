package user

import "github.com/gin-gonic/gin"

type UserController struct {
}

// (con UserController) means mount the method to the struct
func (con UserController) Create(ctx *gin.Context) {
	ctx.JSON(200, "create succeed!")
}

func (con UserController) Index(ctx *gin.Context) {
	ctx.JSON(200, "get succeed!")
}

func (con UserController) Edit(ctx *gin.Context) {
	ctx.JSON(200, "edit succeed!")
}
