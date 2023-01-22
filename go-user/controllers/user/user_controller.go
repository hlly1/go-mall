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
	username, _ := ctx.Cookie("username")
	// username,_ := ctx.Get("username")
	// name, _ := username.(string)
	ctx.JSON(200, "[succeed] username is "+username)
}

func (con UserController) Edit(ctx *gin.Context) {
	ctx.JSON(200, "edit succeed!")
}

func (con UserController) Login(ctx *gin.Context){
	//get paras...
	//check validation code...
	//get authentication info from database...
	//set cookie, 
	//3600s for expire time, 
	//path "/" is default for it is valid globally
	//domain means domain scope, e.g. example.com, should be set when it is deployed
	//(tip)
	// assume a.example.com, b.example.com, example.com shares cookie for all domains
	// domain=a.example.com, cookie will work at this domain only
	//secure, true for https only, false for http and https
	//httpOnly, true for defense xss attack(i.e. js, applet cannot get cookie)
	ctx.SetCookie("username", "alice", 3600,"/","localhost", false, true)
}

func (con UserController) Logout(ctx *gin.Context){
	//Maxage = -1 means delete this cookie
	ctx.SetCookie("username", "alice", -1,"/","localhost", false, true)
}